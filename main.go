package main

import (
	"fmt"
	"log"
	"net/http"
	"todoList/database"
	"todoList/handler"
)

// @title Todo List API
// @version 1.0
// @description With this API you can save Todos.
// @contact.name Alexander Kuepper
// @contact.email kuepper.alexander@web.de
// @host localhost:8080
func main() {
	repositoryTodo := database.NewTodoRepository()
	fmt.Println("REST server is listening on port 8080")
	err := http.ListenAndServe(":8080", handler.NewRouter(&repositoryTodo))
	if err != nil {
		log.Fatal(err)
	}
}
