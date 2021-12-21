package service

import (
	"halill/repository/todo"
	"halill/repository/user"
)

type TodoService interface {
	GetAllTodos()
	GetTodo()
	CreateTodo()
	CompleteTodo()
	DeleteTodo()
}

type todoServiceImpl struct {
	tr todo.TodoRepository
	ur user.UserRepository
}

func (s *todoServiceImpl) GetAllTodos(userId int) {
	todos := s.tr.GetAllByUserId(userId)
	return todos
}

func (s *todoServiceImpl) GetTodo(todoId int) {
	todo := s.tr.GetById(todoId)
	return todo
}

func (s *todoServiceImpl) CreateTodo() {
	todo := s.tr.Create()
	return todo
}

func (s *todoServiceImpl) CompleteTodo() {
	todo := s.tr.Complete()
	return todo
}

func (s *todoServiceImpl) DeleteTodo() {
	todo := s.tr.Delete()
	return todo
}

func NewTodoSerice() TodoService {
	return &todoServiceImpl{}
}
