package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api/handlers"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/validator"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/db"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/services"
	"github.com/labstack/echo/v4"
)

type API struct {
	echo   *echo.Echo
	config config.Config
	logger *logging.Wrapper
	wg     sync.WaitGroup
}

func New(config config.Config, logger *logging.Wrapper, store *db.DB) *API {
	api := &API{
		echo:   echo.New(),
		config: config,
		logger: logger,
	}

	// Configure the echo instance
	api.echo.Debug = api.config.App.Debug
	api.echo.HideBanner = true
	api.echo.HidePort = true
	api.echo.Validator = validator.New()
	api.echo.HTTPErrorHandler = handleErrors(api.logger)
	api.echo.Server.ReadTimeout = api.config.Server.ReadTimeout
	api.echo.Server.WriteTimeout = api.config.Server.WriteTimeout

	// Intialize services
	svc := services.New(store)

	// Register the routes
	api.routes(handlers.New(svc, api.config, api.logger))

	return api
}

func (api *API) Start() error {
	return api.serveHTTP()
}

func handleErrors(logger *logging.Wrapper) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		code := http.StatusInternalServerError
		var message any = "The server encountered a problem and could not process your request."
		if httpError, ok := err.(*echo.HTTPError); ok {
			code = httpError.Code
			switch code {
			case http.StatusNotFound:
				message = "The requested resource could not be found."
			case http.StatusMethodNotAllowed:
				message = fmt.Sprintf("The %s method is not supported for this resource.", c.Request().Method)
			case http.StatusBadRequest:
				message = "The request could not be understood or was missing required parameters."
			case http.StatusInternalServerError:
				message = "The server encountered a problem and could not process your request."
			case http.StatusUnprocessableEntity:
				message = "The request could not be processed due to invalid input."
			default:
				message = httpError.Message
			}
		} else {
			logger.Error(err.Error())
		}
		if err := c.JSON(code, echo.Map{"message": message}); err != nil {
			logger.Error(err.Error())
		}
	}
}
