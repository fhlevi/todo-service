package services

import (
	"fmt"

	"todo-service/database"
	"todo-service/models"
)

// GetTodos retrieves all todo items from the repository.
func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := database.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

// GetTodoByID retrieves a todo item by its ID from the repository.
func GetTodoByID(todoId int) (*models.Todo, error) {
	var todo models.Todo
	result := database.DB.First(&todo, todoId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// CreateTodo adds a new todo item to the repository.
func CreateTodo(todo models.Todo) (models.Todo, error) {
	result := database.DB.Create(&todo)
	if result.Error != nil {
		return models.Todo{}, result.Error
	}
	return todo, nil
}

// UpdateTodo modifies an existing todo item in the repository.
func UpdateTodo(id int, todo models.Todo) (models.Todo, error) {
	var existingTodo models.Todo
	result := database.DB.First(&existingTodo, id)
	if result.Error != nil {
		return models.Todo{}, result.Error
	}

	existingTodo.Todo = todo.Todo
	existingTodo.Date = todo.Date

	result = database.DB.Save(&existingTodo)
	if result.Error != nil {
		return models.Todo{}, result.Error
	}

	return existingTodo, nil
}

// DeleteTodo removes a todo item by its ID from the repository.
func DeleteTodo(id int) error {
	result := database.DB.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}
