package repository

import (
	"errors"
	"todo-service/model"
)

type TodoRepository struct {
	todos []model.Todo
}

func NewTodoRepository() Repository {
	return &TodoRepository{}
}

func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	reversedTodos := make([]model.Todo, len(r.todos))
    for i, item := range r.todos {
        reversedTodos[len(r.todos)-1-i] = item
    }
    return reversedTodos, nil
}

func (r *TodoRepository) GetByID(id string) (model.Todo, error) {
	for _, item := range r.todos {
		if item.ID == id {
			return item, nil
		}
	}
	return model.Todo{}, errors.New("todo not found")
}

func (r *TodoRepository) Create(todo model.Todo) (model.Todo, error) {
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *TodoRepository) Update(id string, todo model.Todo) (model.Todo, error) {
	for i, item := range r.todos {
		if item.ID == id {
			r.todos[i].Todo = todo.Todo
			return r.todos[i], nil
		}
	}
	return model.Todo{}, errors.New("todo not found")
}

func (r *TodoRepository) Delete(id string) error {
	for i, item := range r.todos {
		if item.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
