package user

import (
	"desafio-picpay-go2/internal/common/dto"
	"desafio-picpay-go2/pkg/httputil"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	service   UserService
	secretKey string
	validator *validator.Validate
}

func NewHandler(svc UserService, secretKey string, v *validator.Validate) *handler {
	return &handler{
		service:   svc,
		secretKey: secretKey,
		validator: v,
	}
}
func (h handler) RegisterUserEndpoints(api *echo.Group) {
	//m := middleware.NewWithAuth(h.secretKey)
	api.POST("/users", h.createUser)
	api.POST("/auth/login", h.login)
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

	_, err = h.service.Register(e.Request().Context(), createUserDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.NoContent(http.StatusCreated)
}

func (h handler) login(e echo.Context) error {
	var loginReq dto.LoginRequest
	err := httputil.ReadRequestBody(e.Response(), e.Request(), &loginReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.validator.Struct(loginReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	resp, err := h.service.Login(e.Request().Context(), loginReq)
	if err != nil {
		//TODO: Colocar erro correto
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, resp)
}
