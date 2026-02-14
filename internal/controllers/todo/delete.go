package todo

import (
	"todo-api/internal/pkg/utils"

	"github.com/go-fuego/fuego"
)

// Delete by id
func (c *TodoController) DeleteTodo(ctx fuego.ContextNoBody) (any, error) {
	id := c.GetID(ctx)

	err := c.todoService.DeleteTodo(id)
	if err != nil {
		return nil, c.HandleError(utils.NewNotFoundError("Todo not found", err))
	}

	return c.SuccessResponse("Todo deleted successfully"), nil
}

func (c *TodoController) DeleteAllCompleted(ctx fuego.ContextNoBody) (map[string]interface{}, error) {
	count, err := c.todoService.DeleteAllCompleted()
	if err != nil {
		return nil, c.HandleError(utils.NewInternalServerError("Failed to delete completed todos", err))
	}

	return map[string]interface{}{
		"message":       "Completed todos deleted",
		"deleted_count": count,
	}, nil
}
