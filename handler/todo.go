package handler

import (
	"halill/dto"
	"halill/repository"
	"halill/security"
	"halill/service"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var TodoSet = wire.NewSet(NewTodoHandler, service.NewTodoService, repository.NewUserRepository, repository.NewTodoRepository)

type TodoHandler struct {
	ts service.TodoService
}

func NewTodoHandler(e *echo.Group, ts service.TodoService, jwtSecret string) *TodoHandler {
	handler := &TodoHandler{
		ts: ts,
	}
	config := middleware.JWTConfig{
		Claims:     &security.JwtCustomClaims{},
		SigningKey: []byte(jwtSecret),
	}
	e.Use(middleware.JWTWithConfig(config))
	e.GET("", handler.GetAllTodos)
	e.GET("/:todo_id", handler.GetTodo)
	e.POST("", handler.CreateTodo)
	e.PATCH("/:todo_id", handler.CompleteTodo)
	e.DELETE("/:todo_id", handler.DeleteTodo)

	return handler
}

func (h *TodoHandler) GetAllTodos(c echo.Context) error {
	email := c.Get("user").(*jwt.Token).
		Claims.(*security.JwtCustomClaims).Email
	log.Println(email)
	todos, err := h.ts.GetAllTodos(email)
	if err != nil {
		return err
	}

	return c.JSON(200, todos)
}

func (h *TodoHandler) GetTodo(c echo.Context) error {
	email := c.Get("user").(*jwt.Token).
		Claims.(*security.JwtCustomClaims).Email
	todoID, err := strconv.ParseInt(c.Param("todo_id"), 10, 64)
	if err != nil {
		return err
	}

	todo, err := h.ts.GetTodo(todoID, email)
	if err != nil {
		return err
	}

	return c.JSON(200, todo)
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	email := c.Get("user").(*jwt.Token).
		Claims.(*security.JwtCustomClaims).Email
	request := &dto.CreateTodoRequest{}
	err := c.Bind(request)
	if err != nil {
		return err
	}

	todo, err := h.ts.CreateTodo(request, email)
	if err != nil {
		return err
	}

	return c.JSON(200, todo)
}

func (h *TodoHandler) CompleteTodo(c echo.Context) error {
	email := c.Get("user").(*jwt.Token).
		Claims.(*security.JwtCustomClaims).Email
	todoID, err := strconv.ParseInt(c.Param("todo_id"), 10, 64)
	if err != nil {
		return err
	}

	todo, err := h.ts.CompleteTodo(todoID, email)
	if err != nil {
		return err
	}

	return c.JSON(200, todo)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	email := c.Get("user").(*jwt.Token).
		Claims.(*security.JwtCustomClaims).Email
	todoID, err := strconv.ParseInt(c.Param("todo_id"), 10, 64)
	if err != nil {
		return err
	}

	todo, err := h.ts.DeleteTodo(todoID, email)
	if err != nil {
		return err
	}

	return c.JSON(200, todo)
}
