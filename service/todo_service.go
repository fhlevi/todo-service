package service

import (
	"todo-service/model"
	"todo-service/repository"
)

// Service interface defines the contract for the Todo service
type Service interface {
	GetTodos() ([]model.Todo, error)
	GetTodoByID(id int) (model.Todo, error)
	CreateTodo(todo model.Todo) (model.Todo, error)
	UpdateTodo(id int, todo model.Todo) (model.Todo, error)
	DeleteTodo(id int) error
}

// todoService implements the Service interface for managing Todo items.
type todoService struct {
	repo repository.Repository
}

// NewTodoService creates a new instance of the Todo service.
func NewTodoService(repo repository.Repository) Service {
	return &todoService{repo: repo}
}

// GetTodos retrieves all todo items from the repository.
func (s *todoService) GetTodos() ([]model.Todo, error) {
	return s.repo.GetAll()
}

// GetTodoByID retrieves a todo item by its ID from the repository.
func (s *todoService) GetTodoByID(id int) (model.Todo, error) {
	return s.repo.GetByID(id)
}

// CreateTodo adds a new todo item to the repository.
func (s *todoService) CreateTodo(todo model.Todo) (model.Todo, error) {
	return s.repo.Create(todo)
}

// UpdateTodo modifies an existing todo item in the repository.
func (s *todoService) UpdateTodo(id int, todo model.Todo) (model.Todo, error) {
	return s.repo.Update(id, todo)
}

// DeleteTodo removes a todo item by its ID from the repository.
func (s *todoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}
