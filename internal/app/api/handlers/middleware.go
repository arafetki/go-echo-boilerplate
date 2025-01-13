package handlers

import (
	"net/http"
	"strings"

	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
	"github.com/arafetki/go-echo-boilerplate/internal/jwt"
	"github.com/labstack/echo/v4"
)

var anonymousUser = &sqlc.User{}

func (h *Handler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")
		authorizationHeader := c.Request().Header.Get("Authorization")

		if authorizationHeader == "" {
			c.Set("user", anonymousUser)
			return next(c)
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing authentication token.")
		}

		tokenString := headerParts[1]
		token, err := jwt.HMACCheck(tokenString, h.config.JWT.Key)
		if err != nil {
			h.logger.Error(err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing authentication token.")
		}

		userId, err := token.Claims.GetSubject()
		if err != nil {
			h.logger.Error(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		c.Set("user", &sqlc.User{ID: userId, IsEmailVerified: true})
		return next(c)
	}
}
func (h *Handler) RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*sqlc.User)
		if !ok || user == anonymousUser {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must be authenticated to access this resource.")
		}
		return next(c)
	}
}

func (h *Handler) RequireVerifiedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*sqlc.User)
		if !user.IsEmailVerified {
			return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to access this resource.")
		}
		return next(c)
	}
}
