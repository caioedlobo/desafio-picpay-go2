package user

import (
	"context"
	"desafio-picpay-go2/internal/common"
	"desafio-picpay-go2/internal/domain/user/usecase"
	"desafio-picpay-go2/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	uc usecase.UserUseCase
}

func NewHandler() *handler {
	return &handler{}
}
func (h handler) RegisterUserEndpoints(echo *echo.Echo) {
	echo.POST("/users", h.createUser)
}

func (h handler) createUser(c echo.Context) error {
	var createUserDTO dto.CreateUserRequest
	err := common.ReadJSON(c, &createUserDTO)

	response, err := h.uc.Execute(context.Background(), createUserDTO)
	if err != nil {

	}
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
