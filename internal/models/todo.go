package models

import (
	"time"
)

// Main Todo model
type Todo struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required,min=1,max=200"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type counter struct {
	Total int `json:"total"`
}

// Request DTOs

// Valid data for creating a new Todo
type CreateTodoRequest struct {
	Title string `json:"title" validate:"required,min=1,max=200"`
}

// Valid data for updating an existing Todo
type UpdateTodoRequest struct {
	Title     *string `json:"title" validate:"omitempty,min=1,max=200"`
	Completed *bool   `json:"completed"`
}

// Response DTOs

// Struct to send Todo data to client
type TodoResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Send state of batch operations
type BatchResponse struct {
	Message      string   `json:"message"`
	SuccessCount int      `json:"success_count"`
	FailedIDs    []string `json:"failed_ids,omitempty"`
}

// Convert Todo model to TodoResponse
func (t *Todo) ToResponse() TodoResponse {
	return TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// Contstructor
func NewTodo(title string) *Todo {
	now := time.Now()
	return &Todo{
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
