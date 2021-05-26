package handler

import (
	"todolist/database"
	_ "todolist/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(repositoryTodo *database.TodoRepository) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todo", createTodo(repositoryTodo)).Methods("POST")
	r.HandleFunc("/todo", getAllTodos(repositoryTodo)).Methods("GET")
	r.HandleFunc("/todo/{id}", getTodoById(repositoryTodo)).Methods("GET")
	r.HandleFunc("/todo/{id}", deleteTodoById(repositoryTodo)).Methods("DELETE")
	r.HandleFunc("/todo/{id}", setTodoDone(repositoryTodo)).Methods("PATCH")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}
