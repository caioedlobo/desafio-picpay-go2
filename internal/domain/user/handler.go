package user

import (
	"context"
	"desafio-picpay-go2/internal/common"
	"desafio-picpay-go2/internal/common/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(svc Service) *handler {
	return &handler{
		service: svc,
	}
}
func (h handler) RegisterUserEndpoints(echo *echo.Echo) {
	echo.POST("/users", h.createUser)
}

func (h handler) createUser(c echo.Context) error {
	var createUserDTO dto.CreateUserRequest
	err := common.ReadJSON(c, &createUserDTO)

	_, err = h.service.Save(context.Background(), createUserDTO)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
