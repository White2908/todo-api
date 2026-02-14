package routes

import (
	"github.com/go-fuego/fuego"

	"todo-api/internal/controllers/health"
	"todo-api/internal/controllers/todo"
)

type Routes struct {
	todoController   *todo.TodoController
	healthController *health.HealthController
}

func NewRoutes(
	todoController *todo.TodoController,
	healthController *health.HealthController,
) *Routes {
	return &Routes{
		todoController:   todoController,
		healthController: healthController,
	}
}

func (r *Routes) Register(s *fuego.Server) {

	// Health routes
	healthGroup := fuego.Group(s, "/Health")

	fuego.Get(healthGroup, "/", r.healthController.HealthCheck)
	fuego.Get(healthGroup, "/ready", r.healthController.ReadinessCheck)
	fuego.Get(healthGroup, "/live", r.healthController.LivenessCheck)

	todoGroup := fuego.Group(s, "/features")

	// GET
	fuego.Get(todoGroup, "/", r.todoController.GetAllTodos)
	fuego.Get(todoGroup, "/{id}", r.todoController.GetTodo)
	fuego.Get(todoGroup, "/status/{completed}", r.todoController.GetTodosByStatus)
	fuego.Get(todoGroup, "/search", r.todoController.SearchTodos)

	// POST
	fuego.Post(todoGroup, "/", r.todoController.CreateTodo)

	// PUT/PATCH
	fuego.Put(todoGroup, "/{id}", r.todoController.UpdateTodo)
	fuego.Patch(todoGroup, "/{id}/toggle", r.todoController.ToggleTodoStatus)
	fuego.Patch(todoGroup, "/mark-all-completed", r.todoController.MarkAllCompleted)

	// DELETE
	fuego.Delete(todoGroup, "/{id}", r.todoController.DeleteTodo)
	fuego.Delete(todoGroup, "/completed", r.todoController.DeleteAllCompleted)

}
