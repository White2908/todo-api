package todo

import (
	"todo-api/internal/models"
	"todo-api/internal/pkg/utils"

	"github.com/go-fuego/fuego"
)

// Update complete
func (c *TodoController) UpdateTodo(ctx fuego.ContextWithBody[models.UpdateTodoRequest]) (models.TodoResponse, error) {
	id := ctx.PathParam("id") // <-- changed line

	body, err := ctx.Body()
	if err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewBadRequestError("Invalid request body", err))
	}

	if err := utils.ValidateStruct(body); err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewBadRequestError("Validation failed", err))
	}

	todo, err := c.todoService.UpdateTodo(id, body)
	if err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewNotFoundError("Todo not found", err))
	}

	return *todo, nil
}

// toggle status
func (c *TodoController) ToggleTodoStatus(ctx fuego.ContextNoBody) (models.TodoResponse, error) {
	id := c.GetID(ctx)

	todo, err := c.todoService.ToggleTodoStatus(id)
	if err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewNotFoundError("Todo not found", err))
	}

	return *todo, nil
}

// mark all completed
func (c *TodoController) MarkAllCompleted(ctx fuego.ContextNoBody) (map[string]interface{}, error) {
	count, err := c.todoService.MarkAllCompleted()
	if err != nil {
		return nil, c.HandleError(utils.NewInternalServerError("Failed to mark todos as completed", err))
	}

	return map[string]interface{}{
		"message": "Todos marked as completed",
		"count":   count,
	}, nil
}
