package repository

import "halill/ent"

type TodoRepository interface {
}

type todoRepositoryImpl struct {
	Client *ent.TodoClient
}

func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{}
}

func (r *TodoRepository) FindAllByUserId(userId int) {
	
}
