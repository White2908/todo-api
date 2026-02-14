package todo

import (
	"todo-api/internal/controllers/base"
	"todo-api/internal/services"
)

// handle todo requests
type TodoController struct {
	*base.BaseController
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) *TodoController {
	return &TodoController{
		BaseController: base.NewBaseController(),
		todoService:    todoService,
	}
}
