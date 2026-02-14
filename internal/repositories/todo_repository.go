package repositories

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"todo-api/internal/models"
)

// Repository interface
type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetByID(id string) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(id string, todo *models.Todo) error
	Delete(id string) error
	DeleteAll(ids []string) error
	GetByStatus(completed bool) ([]models.Todo, error)
	Search(query string) ([]models.Todo, error)
}

// From "White" I dont really know how to combine with database currently, so I decided to implement an in-memory repository for purposes(actually because Im stupid).
type InMemoryTodoRepository struct {
	todos map[string]models.Todo
	mu    sync.RWMutex // revent race condition(concurrent access)
}

// constructor
func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string]models.Todo),
	}
}

// CRUD operations
// Get all todos
func (r *InMemoryTodoRepository) GetAll() ([]models.Todo, error) {
	r.mu.RLock()         //read lock
	defer r.mu.RUnlock() //unlock after function ends

	todos := make([]models.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// Get todo by ID
func (r *InMemoryTodoRepository) GetByID(id string) (*models.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todo, exists := r.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo with ID %s not found", id)
	}
	return &todo, nil
}

// Create new todo
func (r *InMemoryTodoRepository) Create(todo *models.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.ID = fmt.Sprintf("%d", len(r.todos)+1)
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	r.todos[todo.ID] = *todo
	return nil
}

// Update existing todo
func (r *InMemoryTodoRepository) Update(id string, todo *models.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.todos[id]
	if !exists {
		return fmt.Errorf("todo with ID %s not found", id)
	}

	todo.ID = id
	todo.CreatedAt = existing.CreatedAt
	todo.UpdatedAt = time.Now()

	r.todos[id] = *todo
	return nil
}

// Delete todo by ID
func (r *InMemoryTodoRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[id]; !exists {
		return fmt.Errorf("todo with ID %s not found", id)
	}

	delete(r.todos, id)
	return nil
}

// Delete multiple todos by IDs
func (r *InMemoryTodoRepository) DeleteAll(ids []string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, id := range ids {
		delete(r.todos, id)
	}
	return nil
}

// Get todos by completion status
func (r *InMemoryTodoRepository) GetByStatus(completed bool) ([]models.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []models.Todo
	for _, todo := range r.todos {
		if todo.Completed == completed {
			result = append(result, todo)
		}
	}
	return result, nil
}

// Search todos by title
func (r *InMemoryTodoRepository) Search(query string) ([]models.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []models.Todo
	for _, todo := range r.todos {
		if strings.Contains(strings.ToLower(todo.Title), strings.ToLower(query)) {
			result = append(result, todo)
		}
	}
	return result, nil
}

// search string
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
