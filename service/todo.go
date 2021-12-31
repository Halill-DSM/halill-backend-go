package service

import (
	"halill/dto"
	"halill/ent"
	"halill/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoService interface {
	GetAllTodos(string) ([]*dto.TodoResponse, error)
	GetTodo(int64, string) (*dto.TodoResponse, error)
	CreateTodo(*dto.CreateTodoRequest, string) (*dto.TodoResponse, error)
	CompleteTodo(int64, string) (*dto.TodoResponse, error)
	DeleteTodo(int64, string) (*dto.TodoResponse, error)
}

type todoServiceImpl struct {
	tr repository.TodoRepository
}

func NewTodoService(tr repository.TodoRepository) TodoService {
	return &todoServiceImpl{
		tr: tr,
	}
}

func (s *todoServiceImpl) GetAllTodos(email string) ([]*dto.TodoResponse, error) {
	todos, err := s.tr.GetAllByEmail(email)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.TodoResponse, 0)
	for _, todo := range todos {
		response = append(response, dto.TodoToDTO(todo))
	}

	return response, nil
}

func (s *todoServiceImpl) GetTodo(todoID int64, email string) (*dto.TodoResponse, error) {
	todo, err := s.tr.Get(todoID)
	if err != nil {
		return nil, err
	}

	if todo.Edges.User.ID != email {
		return nil, echo.NewHTTPError(http.StatusForbidden, "해당 요청에 대한 권한이 없습니다.")
	}

	return dto.TodoToDTO(todo), nil
}

func (s *todoServiceImpl) CreateTodo(request *dto.CreateTodoRequest, email string) (*dto.TodoResponse, error) {
	todo := &ent.Todo{
		Title:    request.Title,
		Content:  request.Content,
		Deadline: request.Deadline,
		Edges: ent.TodoEdges{
			User: &ent.User{ID: email},
		},
	}

	newTodo, err := s.tr.Create(todo)
	if err != nil {
		return nil, err
	}
	return dto.TodoToDTO(newTodo), nil
}

func (s *todoServiceImpl) CompleteTodo(todoID int64, email string) (*dto.TodoResponse, error) {
	todo, err := s.tr.Get(todoID)
	if err != nil {
		return nil, err
	}

	if todo.Edges.User.ID != email {
		return nil, echo.NewHTTPError(http.StatusForbidden, "해당 요청에 대한 권한이 없습니다.")
	}

	_, err = s.tr.Complete(todoID)
	if err != nil {
		return nil, err
	}

	return dto.TodoToDTO(todo), nil
}

func (s *todoServiceImpl) DeleteTodo(todoID int64, email string) (*dto.TodoResponse, error) {
	todo, err := s.tr.Get(todoID)
	if err != nil {
		return nil, err
	}

	if todo.Edges.User.ID != email {
		return nil, echo.NewHTTPError(http.StatusForbidden, "해당 요청에 대한 권한이 없습니다.")
	}

	_, err = s.tr.Delete(todoID)
	if err != nil {
		return nil, err
	}

	return dto.TodoToDTO(todo), nil
}
