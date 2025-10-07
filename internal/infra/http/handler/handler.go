package handler

import (
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/infra/container"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	echo *echo.Echo
}

func RegisterHandler(echo *echo.Echo, c *container.Container) {
	user.NewHandler(c.UserService).RegisterUserEndpoints(echo)
}
