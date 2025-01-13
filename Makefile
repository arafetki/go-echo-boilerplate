## Colors
COLOR_RESET   = \033[0m
COLOR_INFO    = \033[32m
COLOR_COMMENT = \033[33m

MAIN_PACKAGE_PATH = ./cmd/api
BINARY_NAME = api

.PHONY: help
## Help
help:
	@printf "${COLOR_COMMENT}Usage:${COLOR_RESET}\n"
	@printf " make [target] [args...]\n\n"
	@printf "${COLOR_COMMENT}Available targets:${COLOR_RESET}\n"
	@awk '/^[a-zA-Z\-\0-9\.@]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf " ${COLOR_INFO}%-16s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: tidy
## Format code and tidy modfile
tidy:
	go mod tidy -v
	go fmt ./...


.PHONY: audit
## Run quality control checks
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


.PHONY: test
## Run unit tests
test:
	go test -race -buildvcs -vet=off ./...


.PHONY: build
## Build the application
build:
	go build -o=./build/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run
## Run the application
run: build
	./build/${BINARY_NAME}

.PHONY: run-dev
## Run the application with reloading on file changes
run-dev:
	go run github.com/air-verse/air@latest \
		--build.cmd "go build -o=./tmp/${BINARY_NAME} ${MAIN_PACKAGE_PATH}" --build.bin "./tmp/${BINARY_NAME}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, sql, tmpl, html, css, scss, js, mjs, cjs, json, yaml, yml, toml, ini" \
		--misc.clean_on_exit "true"

.PHONY: migrations/new
## create a new database migration
migrations/new:
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -seq -ext=.sql -dir=./assets/migrations ${name}


.PHONY: migrations/up
## apply all up database migrations
migrations/up:
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="postgresql://${DATABASE_DSN}" up


.PHONY: migrations/down
## apply all down database migrations
migrations/down:
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="postgresql://${DATABASE_DSN}" down