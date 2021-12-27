package handler

import (
	"halill/dto"
	"halill/repository"
	"halill/security"
	"halill/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var UserSet = wire.NewSet(NewUserHandler, service.NewUserSerice, repository.NewUserRepository, security.NewJWTProvider)

type UserHandler struct {
	us service.UserService
}

func NewUserHandler(e *echo.Group, us service.UserService) *UserHandler {
	handler := &UserHandler{
		us: us,
	}
	e.POST("/login", handler.Login)
	e.POST("/signup", handler.Register)
	e.PUT("/login", handler.Refresh)
	return handler
}

func (h *UserHandler) Login(c echo.Context) error {
	request := &dto.LoginRequest{}
	err := c.Bind(request)
	if err != nil {
		return err
	}

	token, err := h.us.LoginUser(request)
	if err != nil {
		return err
	}

	return c.JSON(200, token)
}

func (h *UserHandler) Register(c echo.Context) error {
	request := &dto.RegistRequest{}
	err := c.Bind(request)
	if err != nil {
		return err
	}

	todo, err := h.us.RegistUser(request)
	if err != nil {
		return err
	}

	return c.JSON(200, todo)
}

func (h *UserHandler) Refresh(c echo.Context) error {
	request := &dto.RefreshTokenRequest{}
	err := c.Bind(request)
	if err != nil {
		return err
	}

	response, err := h.us.RefreshToken(request)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}
