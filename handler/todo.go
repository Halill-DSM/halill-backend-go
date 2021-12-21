package handler

import (
	"halill/service"

	"github.com/labstack/echo"
)

type TodoHandler struct {
	ts service.TodoService
}

func NewTodoHandler(e *echo.Echo, ts service.TodoService) {
	handler := &TodoHandler{
		ts: ts,
	}
	e.GET("/todo", handler.GetAllTodos)
	e.GET("/todo/:todo_id", handler.GetTodo)
	e.POST("/todo", handler.CreateTodo)
	e.PATCH("/todo/:todo_id", handler.CompleteTodo)
	e.DELETE("/todo/:todo_id", handler.DeleteTodo)
}

func (h *TodoHandler) GetAllTodos(c echo.Context) {
	todos := h.ts.GetAllTodos()
	c.JSON(200, todos)
}

func (h *TodoHandler) GetTodo(c echo.Context) {
	todo := h.ts.GetTodo(c.Param("todo_id"))
	c.JSON(200, todo)
}

func (h *TodoHandler) CreateTodo(c echo.Context) {
	todo := h.ts.CreateTodo()
	c.JSON(200, todo)
}

func (h *TodoHandler) CompleteTodo(c echo.Context) {

	todo := h.ts.CompleteTodo()
	c.JSON(200, todo)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) {
	todo := h.ts.DeleteTodo()
	c.JSON(200, todo)
}
