package router

import (
	"github.com/gorilla/mux"
	"todo-service/controller"
    httpSwagger "github.com/swaggo/http-swagger"
)

func Router(todoController *controller.TodoController) *mux.Router {
	router := mux.NewRouter();

	// Swagger endpoint
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := router.PathPrefix("/api/v1").Subrouter();
	api.HandleFunc("/todos", todoController.GetTodos).Methods("GET")
	api.HandleFunc("/todos/{id}", todoController.GetTodoByID).Methods("GET")
	api.HandleFunc("/todos", todoController.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", todoController.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", todoController.DeleteTodo).Methods("DELETE")

	return router
}
