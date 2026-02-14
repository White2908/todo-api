package services

import (
	"todo-api/internal/models"
	"todo-api/internal/repositories"
)

// Service interface
type TodoService interface {
	GetAllTodos() ([]models.TodoResponse, error)
	GetTodo(id string) (*models.TodoResponse, error)
	GetTodosByStatus(completed bool) ([]models.TodoResponse, error)
	SearchTodos(query string) ([]models.TodoResponse, error)
	CreateTodo(req models.CreateTodoRequest) (*models.TodoResponse, error)
	CreateBatchTodos(reqs []models.CreateTodoRequest) ([]models.TodoResponse, error)
	UpdateTodo(id string, req models.UpdateTodoRequest) (*models.TodoResponse, error)
	ToggleTodoStatus(id string) (*models.TodoResponse, error)
	MarkAllCompleted() (int, error)
	DeleteTodo(id string) error
	DeleteBatchTodos(ids []string) (int, []string, error)
	DeleteAllCompleted() (int, error)
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

// Service operations

// READ operations
func (s *todoService) GetAllTodos() ([]models.TodoResponse, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = todo.ToResponse()
	}
	return responses, nil
}

func (s *todoService) GetTodo(id string) (*models.TodoResponse, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := todo.ToResponse()
	return &response, nil
}

func (s *todoService) GetTodosByStatus(completed bool) ([]models.TodoResponse, error) {
	todos, err := s.repo.GetByStatus(completed)
	if err != nil {
		return nil, err
	}

	responses := make([]models.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = todo.ToResponse()
	}
	return responses, nil
}

func (s *todoService) SearchTodos(query string) ([]models.TodoResponse, error) {
	todos, err := s.repo.Search(query)
	if err != nil {
		return nil, err
	}

	responses := make([]models.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = todo.ToResponse()
	}
	return responses, nil
}

// CREATE operations
func (s *todoService) CreateTodo(req models.CreateTodoRequest) (*models.TodoResponse, error) {

	todo := models.NewTodo(req.Title)

	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}

	response := todo.ToResponse()
	return &response, nil
}

func (s *todoService) CreateBatchTodos(reqs []models.CreateTodoRequest) ([]models.TodoResponse, error) {
	var responses []models.TodoResponse

	for _, req := range reqs {
		todo, err := s.CreateTodo(req)
		if err != nil {
			continue
		}
		responses = append(responses, *todo)
	}

	return responses, nil
}

// UPDATE operations
func (s *todoService) UpdateTodo(id string, req models.UpdateTodoRequest) (*models.TodoResponse, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		existing.Title = *req.Title
	}
	if req.Completed != nil {
		existing.Completed = *req.Completed
	}

	if err := s.repo.Update(id, existing); err != nil {
		return nil, err
	}

	response := existing.ToResponse()
	return &response, nil
}

func (s *todoService) ToggleTodoStatus(id string) (*models.TodoResponse, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	existing.Completed = !existing.Completed

	if err := s.repo.Update(id, existing); err != nil {
		return nil, err
	}

	response := existing.ToResponse()
	return &response, nil
}

func (s *todoService) MarkAllCompleted() (int, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	count := 0
	for _, todo := range todos {
		if !todo.Completed {
			todo.Completed = true
			if err := s.repo.Update(todo.ID, &todo); err == nil {
				count++
			}
		}
	}

	return count, nil
}

// DELETE operations
func (s *todoService) DeleteTodo(id string) error {
	return s.repo.Delete(id)
}

func (s *todoService) DeleteBatchTodos(ids []string) (int, []string, error) {
	var failedIDs []string
	successCount := 0

	for _, id := range ids {
		if err := s.repo.Delete(id); err != nil {
			failedIDs = append(failedIDs, id)
		} else {
			successCount++
		}
	}

	return successCount, failedIDs, nil
}

func (s *todoService) DeleteAllCompleted() (int, error) {
	todos, err := s.repo.GetByStatus(true)
	if err != nil {
		return 0, err
	}

	ids := make([]string, len(todos))
	for i, todo := range todos {
		ids[i] = todo.ID
	}

	if err := s.repo.DeleteAll(ids); err != nil {
		return 0, err
	}

	return len(ids), nil
}
