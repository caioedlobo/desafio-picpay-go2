package handler

import (
	"desafio-picpay-go2/internal/domain/user"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	echo *echo.Echo
}

func RegisterHandler(echo *echo.Echo) {
	user.NewHandler().RegisterUserEndpoints(echo)
}
