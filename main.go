package main

import (
	"log"
	"os"

	"todo-service/database"
	"todo-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.RedirectTrailingSlash = false

	// CORS middleware
	r.Use(func(c *gin.Context) {
		// Set header CORS sebelum request diproses
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Max-Age", "43200")
		
		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "service": "unified-service"})
	})

	// Serve Swagger UI at root docs path
	r.GET("/api/docs", func(c *gin.Context) {
		c.File("./docs/index.html")
	})
	
	// Serve swagger JSON
	r.GET("/docs/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	
	// Serve swagger YAML
	r.GET("/docs/swagger.yaml", func(c *gin.Context) {
		c.File("./docs/swagger.yaml")
	})

	// Todo routes
	todo := r.Group("/api/todo")
	{
		todo.GET("/", handlers.GetTodos)
		todo.POST("/", handlers.CreateTodo)
		todo.GET("/:id", handlers.GetTodoByID)
		todo.PUT("/:id", handlers.UpdateTodo)
		todo.DELETE("/:id", handlers.DeleteTodo)
	}

	port := getEnv("PORT", "8080")
	log.Printf("Unified service running on port %s", port)
	log.Printf("🚀 Swagger UI Documentation: http://localhost:%s/api/docs", port)
	log.Printf("📄 Swagger JSON: http://localhost:%s/docs/swagger.json", port)
	r.Run(":" + port)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}