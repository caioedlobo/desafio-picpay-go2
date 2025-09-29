package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	echo *echo.Echo
}

func NewHandler() *Handler {
	e := echo.New()
	return &Handler{e}
}

func (h Handler) Listen(port string) error {
	h.registerUserEndpoints()
	return h.echo.Start(fmt.Sprintf(":%s", port))
}
