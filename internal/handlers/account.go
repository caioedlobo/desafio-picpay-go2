package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) registerUserEndpoints() {
	h.echo.GET("/users", h.helloWorld)
}

func (h Handler) helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
