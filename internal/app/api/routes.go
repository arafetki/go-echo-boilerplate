package api

import (
	"net/http"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api/handlers"
	"github.com/labstack/echo/v4/middleware"
)

func (api *API) routes(h *handlers.Handler) {

	// Register global middleware
	api.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human} | ${remote_ip} | ${user_agent} | error: ${error}\n",
	}))
	api.echo.Use(middleware.Recover())
	api.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	}))

	// Health check endpoint for load balancers
	api.echo.GET("/health", h.HealthCheckHandler)

}
