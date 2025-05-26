package controller

import (
	"encoding/json"
	"net/http"
	"todo-service/model"
	"todo-service/service"
	"github.com/gorilla/mux"
	"time"
	"github.com/google/uuid"
)

// TodoController handles HTTP requests for todo items.
type TodoController struct {
	service service.Service
}

// NewTodoController creates a new instance of TodoController with the provided service.
func NewTodoController(service service.Service) *TodoController {
	return &TodoController{service: service}
}

// RegisterRoutes registers the routes for the TodoController.
func (c *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.service.GetTodos()
	if err != nil {
		http.Error(w, "Error retrieving todos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// GetTodoByID retrieves a todo item by its ID.
func (c *TodoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo, err := c.service.GetTodoByID(params["id"])
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// CreateTodo creates a new todo item.
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Todo string `json:"todo"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Todo == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	todo := model.Todo{
		ID:   uuid.New().String(), // generate UUID
		Todo: req.Todo,
		Date: time.Now().Format("Monday, January 2, 2006 at 3:04 PM"),
	}

	createdTodo, err := c.service.CreateTodo(todo)
	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdTodo)
}

// UpdateTodo updates an existing todo item.
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	updatedTodo, err := c.service.UpdateTodo(params["id"], todo)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// DeleteTodo deletes a todo item by its ID.
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Retrieve the todo item before deletion
    todo, err := c.service.GetTodoByID(params["id"])
    if err != nil {
        http.Error(w, "Todo not found", http.StatusNotFound)
        return
    }

	// Proceed to delete the todo item
	err = c.service.DeleteTodo(params["id"]);
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Return the deleted todo item as the response
	json.NewEncoder(w).Encode(todo)
}
