package controller

import (
	"encoding/json"
	"net/http"
	"todo-service/model"
	"todo-service/service"
	"github.com/gorilla/mux"
	"time"
	// "github.com/google/uuid"
	"strconv"
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
// GetTodos godoc
// @Summary      Get all todos
// @Description  Retrieve all todo items
// @Tags         todos
// @Produce      json
// @Success      200  {array}  model.Todo
// @Failure      500  {string} string "Error retrieving todos"
// @Router       /api/v1/todos [get]
func (c *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.service.GetTodos()
	if err != nil {
		http.Error(w, "Error retrieving todos: " + err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// GetTodoByID retrieves a todo item by its ID.
// GetTodoByID godoc
// @Summary      Get todo by ID
// @Description  Retrieve a todo item by its ID
// @Tags         todos
// @Produce      json
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  model.Todo
// @Failure      404  {string}  string "Todo not found"
// @Router       /api/v1/todos/{id} [get]
func (c *TodoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := c.service.GetTodoByID(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// CreateTodo creates a new todo item.
// CreateTodo godoc
// @Summary      Create a new todo
// @Description  Create a new todo item
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        request  body   model.TodoRequest  true  "Todo to create"
// @Success      200   {object}  model.Todo
// @Failure      400   {string}  string "Invalid request"
// @Failure      500   {string}  string "Error creating todo"
// @Router       /api/v1/todos [post]
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req model.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Todo == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	todo := model.Todo{
		// ID:   uuid.New().String(), // generate UUID
		Todo: req.Todo,
		Date: time.Now().Format("Monday, January 2, 2006 at 3:04 PM"),
	}

	createdTodo, err := c.service.CreateTodo(todo)
	if err != nil {
		http.Error(w, "Error creating todo: " + err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdTodo)
}

// UpdateTodo updates an existing todo item.
// UpdateTodo godoc
// @Summary      Update a todo
// @Description  Update an existing todo item
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      string      true  "Todo ID"
// @Param        todo  body      model.TodoRequest  true  "Todo to update"
// @Success      200   {object}  model.Todo
// @Failure      404   {string}  string "Todo not found"
// @Router       /api/v1/todos/{id} [put]
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req model.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Todo == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	todo := model.Todo{
		Todo: req.Todo,
		Date: time.Now().Format("Monday, January 2, 2006 at 3:04 PM"), // inilah kuncinya
	}

	updatedTodo, err := c.service.UpdateTodo(id, todo)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// DeleteTodo deletes a todo item by its ID.
// DeleteTodo godoc
// @Summary      Delete a todo
// @Description  Delete a todo item by its ID
// @Tags         todos
// @Produce      json
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  model.Todo
// @Failure      404  {string}  string "Todo not found"
// @Router       /api/v1/todos/{id} [delete]
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve the todo item before deletion
    todo, err := c.service.GetTodoByID(id)
    if err != nil {
        http.Error(w, "Todo not found", http.StatusNotFound)
        return
    }

	// Proceed to delete the todo item
	err = c.service.DeleteTodo(id);
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Return the deleted todo item as the response
	json.NewEncoder(w).Encode(todo)
}
