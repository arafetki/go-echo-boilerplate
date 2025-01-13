package main

import (
	"os"
	"runtime/debug"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/db"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
)

func main() {
	logger := logging.NewSlogLogger(os.Stdout)
	if err := run(logger); err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run(logger *logging.Wrapper) error {
	cfg := config.Load()
	logger.SetLevel(cfg.Logger.Level)

	// Connect to the database
	db, err := db.Pool(cfg.Database.Dsn, cfg.Database.Automigrate)
	if err != nil {
		return err
	}
	defer db.Close()
	logger.Info("Database connection established sucessfully")

	api := api.New(cfg, logger, db)
	return api.Start()
}
