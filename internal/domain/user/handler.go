package user

import (
	"context"
	"desafio-picpay-go2/internal/common/dto"
	"desafio-picpay-go2/pkg/httputil"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(svc Service, v *validator.Validate) *handler {
	return &handler{
		service:   svc,
		validator: v,
	}
}
func (h handler) RegisterUserEndpoints(echo *echo.Echo) {
	echo.POST("/users", h.createUser)
}

func (h handler) createUser(e echo.Context) error {
	var createUserDTO dto.CreateUserRequest
	err := httputil.ReadRequestBody(e.Response(), e.Request(), &createUserDTO)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.validator.Struct(createUserDTO)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	_, err = h.service.Save(context.Background(), createUserDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.NoContent(http.StatusCreated)
}
