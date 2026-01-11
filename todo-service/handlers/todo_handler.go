package handlers

import (
	"net/http"
	"strconv"
	"time"

	"todo-service/models"
	"todo-service/services"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := services.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := services.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var req models.TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	
	reqBody := models.Todo{
		Todo: req.Todo,
		Date: time.Now().Format("Monday, January 2, 2006 at 3:04 PM"),
	}

	todo, err := services.CreateTodo(reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req models.TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	reqBody := models.Todo{
		Todo: req.Todo,
		Date: time.Now().Format("Monday, January 2, 2006 at 3:04 PM"),
	}

	todo, err := services.UpdateTodo(id, reqBody)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Retrieve the todo item before deletion
    todo, err := services.GetTodoByID(id)
    if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

	// Proceed to delete the todo item
	err = services.DeleteTodo(id);
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Return the deleted todo item as the response
	c.JSON(http.StatusOK, todo)
}