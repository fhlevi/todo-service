package repository

import (
	"database/sql"
	"errors"
	"todo-service/model"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) Repository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	rows, err := r.DB.Query("SELECT id, todo, date FROM todo ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Todo, &todo.Date); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) GetByID(id int) (model.Todo, error) {
	var todo model.Todo
	err := r.DB.QueryRow("SELECT id, todo, date FROM playing_with_leapcell WHERE id = $1", id).Scan(&todo.ID, &todo.Todo, &todo.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Todo{}, errors.New("todo not found")
		}
		return model.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) Create(todo model.Todo) (model.Todo, error) {
	err := r.DB.QueryRow(
		"INSERT INTO playing_with_leapcell (todo, date) VALUES ($1, $2) RETURNING id",
		todo.Todo, todo.Date,
	).Scan(&todo.ID)

	if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) Update(id int, todo model.Todo) (model.Todo, error) {
	result, err := r.DB.Exec("UPDATE playing_with_leapcell SET todo = $1, date = $2 WHERE id = $3", todo.Todo, todo.Date, id)
	if err != nil {
		return model.Todo{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Todo{}, err
	}
	if rowsAffected == 0 {
		return model.Todo{}, errors.New("todo not found")
	}
	todo.ID = id
	return todo, nil
}

func (r *TodoRepository) Delete(id int) error {
	result, err := r.DB.Exec("DELETE FROM playing_with_leapcell WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}
