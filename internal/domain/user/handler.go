package user

import (
	"desafio-picpay-go2/internal/common/dto"
	error2 "desafio-picpay-go2/internal/common/error"
	"desafio-picpay-go2/internal/infra/http/middleware"
	"desafio-picpay-go2/pkg/fault"
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
	m := middleware.NewWithAuth(h.secretKey)
	api.POST("/users", h.create)
	api.GET("/users/me", h.get, echo.WrapMiddleware(m.WithAuth))
	api.POST("/auth/login", h.login)
}

func (h handler) create(e echo.Context) error {
	var createUserDTO dto.CreateUserRequest
	err := httputil.ReadRequestBody(e.Response(), e.Request(), &createUserDTO)
	if err != nil {
		fault.NewHTTPError(e.Response(), err)
		return err
	}

	err = h.validator.Struct(createUserDTO)
	if err != nil {
		fault.NewHTTPError(
			e.Response(),
			fault.New(
				"failed to validate body",
				fault.WithValidationError(error2.ToFaultErrors(err))),
		)
		return err
	}

	_, err = h.service.Register(e.Request().Context(), createUserDTO)

	if err != nil {
		fault.NewHTTPError(e.Response(), err)
		return err
	}

	return e.NoContent(http.StatusCreated)
}

func (h handler) get(e echo.Context) error {
	resp, err := h.service.Get(e.Request().Context())
	if err != nil {
		fault.NewHTTPError(e.Response(), err)
		return err
	}
	return e.JSON(http.StatusOK, resp)
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
