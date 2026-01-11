package main

import (
	"log"
	"os"

	"todo-service/database"
	"todo-service/handlers"
	
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init() // Initialize connection

	r := gin.Default()

	todo := r.Group("/api/todo")
	{
		todo.GET("/", handlers.GetTodos)
		todo.GET("/:id", handlers.GetTodoByID)
		todo.POST("/", handlers.CreateTodo)
		todo.PUT("/:id", handlers.UpdateTodo)
		todo.DELETE("/:id", handlers.DeleteTodo)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "service": "todo-service"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Todo service running on port %s", port)
	r.Run(":" + port)
}