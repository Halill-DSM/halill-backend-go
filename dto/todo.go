package dto

import (
	"halill/ent"
	"time"
)

type CreateTodoRequest struct {
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	Deadline *time.Time `json:"deadline,omitempty"`
}

type TodoResponse struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Deadline    *time.Time `json:"deadline"`
	IsCompleted bool       `json:"is_completed"`
}

func TodoToDTO(src *ent.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          src.ID,
		Title:       src.Title,
		Content:     src.Content,
		Deadline:    src.Deadline,
		IsCompleted: src.IsCompleted,
	}
}
