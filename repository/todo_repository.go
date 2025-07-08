package repository

import "todo-service/model"

type Repository interface {
	// GetAll retrieves all todo items.
	GetAll() ([]model.Todo, error)
	// GetByID retrieves a todo item by its ID.
	GetByID(id int) (model.Todo, error)
	// Create adds a new todo item.
	Create(todo model.Todo) (model.Todo, error)
	// Update modifies an existing todo item.
	Update(id int, todo model.Todo) (model.Todo, error)
	// Delete removes a todo item by its ID.
	Delete(id int) error
}