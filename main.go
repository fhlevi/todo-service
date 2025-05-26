package main

import (
	"log"
	"net/http"
	"todo-service/config"
	"todo-service/router"
	"todo-service/controller"
	"todo-service/repository"
	"todo-service/service"
	"github.com/rs/cors"
)

func main() {
	config.Init() // Initialize configuration

	repo := repository.NewTodoRepository()
	service := service.NewTodoService(repo)
	controller := controller.NewTodoController(service)
	r := router.Router(controller)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Server running on port", config.AppPort)
	log.Fatal(http.ListenAndServe(config.AppPort, handler))
}
