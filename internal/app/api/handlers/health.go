package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":    "available",
		"timestamp": time.Now().UnixNano(),
	})
}
