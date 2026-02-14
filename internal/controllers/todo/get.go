package todo

import (
	"strconv"

	"todo-api/internal/models"
	"todo-api/internal/pkg/utils"

	"github.com/go-fuego/fuego"
)

// Get all todos
func (c *TodoController) GetAllTodos(ctx fuego.ContextNoBody) ([]models.TodoResponse, error) {
	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		return nil, c.HandleError(utils.NewInternalServerError("Failed to fetch todos", err))
	}
	return todos, nil
}

// Search by id
func (c *TodoController) GetTodo(ctx fuego.ContextNoBody) (models.TodoResponse, error) {
	id := c.GetID(ctx)

	todo, err := c.todoService.GetTodo(id)
	if err != nil {
		return models.TodoResponse{}, c.HandleError(utils.NewNotFoundError("Todo not found", err))
	}

	return *todo, nil
}

// Search by status
func (c *TodoController) GetTodosByStatus(ctx fuego.ContextNoBody) ([]models.TodoResponse, error) {
	status := ctx.Request().PathValue("completed")
	completed, _ := strconv.ParseBool(status)

	todos, err := c.todoService.GetTodosByStatus(completed)
	if err != nil {
		return nil, c.HandleError(utils.NewInternalServerError("Failed to fetch todos", err))
	}

	return todos, nil
}

// Search
func (c *TodoController) SearchTodos(ctx fuego.ContextNoBody) ([]models.TodoResponse, error) {
	query := ctx.Request().URL.Query().Get("q")
	todos, err := c.todoService.SearchTodos(query)
	if err != nil {
		return nil, c.HandleError(utils.NewInternalServerError("Failed to search todos", err))
	}
	return todos, nil
}
