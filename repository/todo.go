package repository

import (
	"context"
	"halill/ent"
	"halill/ent/todo"
	"halill/ent/user"
)

type TodoRepository interface {
	GetAllByEmail(string) ([]*ent.Todo, error)
	Get(int64) (*ent.Todo, error)
	Create(*ent.Todo) (*ent.Todo, error)
	Complete(int64) (*ent.Todo, error)
	Delete(int64) (*ent.Todo, error)
}

type todoRepositoryImpl struct {
	db *ent.Client
}

func NewTodoRepository(db *ent.Client) TodoRepository {
	return &todoRepositoryImpl{
		db: db,
	}
}

func (r *todoRepositoryImpl) GetAllByEmail(email string) ([]*ent.Todo, error) {
	result := r.db.Todo.Query().
		Where(todo.HasUserWith(user.ID(email))).
		WithUser().
		AllX(context.TODO())

	return result, nil
}

func (r *todoRepositoryImpl) Get(todoID int64) (*ent.Todo, error) {
	return nil, nil
}

func (r *todoRepositoryImpl) Create(*ent.Todo) (*ent.Todo, error) {
	return nil, nil
}

func (r *todoRepositoryImpl) Complete(todoID int64) (*ent.Todo, error) {
	return nil, nil
}

func (r *todoRepositoryImpl) Delete(todoID int64) (*ent.Todo, error) {
	return nil, nil
}
