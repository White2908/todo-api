package todo

import (
	"todo-api/internal/models"
	"todo-api/internal/pkg/utils"

	"github.com/go-fuego/fuego"
)

// POST /todos
func (c *TodoController) CreateTodo(ctx fuego.ContextWithBody[models.CreateTodoRequest]) (models.TodoResponse, error) {
	body, err := ctx.Body()
	if err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewBadRequestError("Invalid request body", err))
	}

	if err := utils.ValidateStruct(body); err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewBadRequestError("Validation failed", err))
	}

	todo, err := c.todoService.CreateTodo(body)
	if err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewInternalServerError("Failed to create todo", err))
	}

	return *todo, nil
}
