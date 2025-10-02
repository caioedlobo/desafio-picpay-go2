package handler

import (
	"desafio-picpay-go2/internal/dto/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) registerUserEndpoints() {
	h.echo.POST("/users", h.createUser)
}

func (h Handler) createUser(c echo.Context) error {
	var createUserDTO user.CreateUserRequest
	err := readJSON(c, &createUserDTO)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
