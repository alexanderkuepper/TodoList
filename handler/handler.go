package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"homework-SoerenDev-391298709521/database"
	"homework-SoerenDev-391298709521/model"
	"net/http"
	"strconv"
)

// @Title Create
// @tags todo
// @Summary Create a todo
// @Description you can create a todo with this method
// @Accept json
// @Produce json
// @Param body body model.PostTodo true "todo"
// @Success 201 "Status Created"
// @Failure 400 "Not Found"
// @Router /todo [post]
func createTodo(repositoryTodo *database.TodoRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := &model.PostTodo{}
		err := json.NewDecoder(r.Body).Decode(todo)
		repositoryTodo.AddTodo(*todo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}

// @Title Get
// @tags todo
// @Summary Show all todos
// @Description get all Todos
// @Accept  json
// @Success 200 {array} model.Todo
// @Failure 500 "Server Error"
// @Router /todo [get]
func getAllTodos(repository *database.TodoRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := repository.GetAllTodos()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		err := json.NewEncoder(w).Encode(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// @Title Get
// @tags todo
// @Summary Show a todo
// @Description get todo by ID
// @Accept  json
// @Param id path int true "ID"
// @Success 200 {object} model.Todo
// @Failure 404 "Not Found"
// @Failure 500 "Server Error"
// @Router /todo/{id} [get]
func getTodoById(repositoryTodo *database.TodoRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		todo, err := repositoryTodo.GetTodoById(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			err = json.NewEncoder(w).Encode(todo)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// @Title Delete
// @tags todo
// @Summary Delete the todo
// @Description everyone can delete the todo
// @Param id path int true "ID"
// @Success 204 "No Content"
// @Failure 404 "Not Found"
// @Router /todo/{id} [delete]
func deleteTodoById(repositoryTodo *database.TodoRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		err := repositoryTodo.DeleteTodo(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

// @Title Patch
// @tags todo
// @Summary Set todo done
// @Description with this method you can set the done bit true
// @Param id path int true "ID"
// @Success 204 "No Content"
// @Failure 404 "Not Found"
// @Router /todo/{id} [patch]
func setTodoDone(repositoryTodo *database.TodoRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		err := repositoryTodo.TodoIsDone(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
