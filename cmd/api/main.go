package main

import (
	"os"

	"github.com/go-fuego/fuego"

	"todo-api/internal/controllers/health"
	"todo-api/internal/controllers/todo"
	"todo-api/internal/repositories"
	"todo-api/internal/routes"
	"todo-api/internal/services"
)

func main() {
	// set port from environment variable (default:8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create Fuego server
	s := fuego.NewServer(
		fuego.WithAddr("localhost:" + port),
	)

	todoRepo := repositories.NewInMemoryTodoRepository()  // Repository save in ram
	todoService := services.NewTodoService(todoRepo)      // Service Todo
	todoController := todo.NewTodoController(todoService) // Controller Todo
	healthController := health.NewHealthController()      // Controller health

	// Register routes
	appRoutes := routes.NewRoutes(todoController, healthController)
	appRoutes.Register(s)

	s.Run()
}
