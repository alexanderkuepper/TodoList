package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoList/database"
	"todoList/model"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var repository database.TodoRepository
var router *mux.Router

func setup() {
	repository = database.NewTodoRepository()
	firstTodo := model.PostTodo{Description: "Write an API", DueDate: "07.05.2021", Priority: 10}
	secondTodo := model.PostTodo{Description: "Get the job", DueDate: "07.05.2021", Priority: 10}
	repository.AddTodo(firstTodo)
	repository.AddTodo(secondTodo)
	router = NewRouter(&repository)
}

func TestGetTodoById(t *testing.T) {
	setup()
	req, _ := http.NewRequest("GET", "/todo/0", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	expected := "{\"Id\":0,\"Description\":\"Write an API\",\"DueDate\":\"07.05.2021\",\"Done\":false,\"Priority\":10}\n"
	assert.Equal(t, expected, rec.Body.String())
}

func TestGetAllTodos(t *testing.T) {
	setup()
	req, _ := http.NewRequest("GET", "/todo", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	expected := "[{\"Id\":0,\"Description\":\"Write an API\",\"DueDate\":\"07.05.2021\",\"Done\":false,\"Priority\":10}," +
		"{\"Id\":1,\"Description\":\"Get the job\",\"DueDate\":\"07.05.2021\",\"Done\":false,\"Priority\":10}]\n"
	assert.Equal(t, expected, rec.Body.String())
}

func TestDeleteTodoById(t *testing.T) {
	setup()
	req, _ := http.NewRequest("DELETE", "/todo/0", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	req, _ = http.NewRequest("GET", "/todo", nil)
	router.ServeHTTP(rec, req)
	expected := "[{\"Id\":1,\"Description\":\"Get the job\",\"DueDate\":\"07.05.2021\",\"Done\":false,\"Priority\":10}]\n"
	assert.Equal(t, expected, rec.Body.String())
}

func TestCreateTodo(t *testing.T) {
	setup()
	var jsonStr = []byte(`{"description": "Have fun","dueDate": "07.05.2021", "priority": 11}`)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)

	req, _ = http.NewRequest("GET", "/todo/2", nil)
	router.ServeHTTP(rec, req)
	expected := "{\"Id\":2,\"Description\":\"Have fun\",\"DueDate\":\"07.05.2021\",\"Done\":false,\"Priority\":11}\n"
	assert.Equal(t, expected, rec.Body.String())
}

func TestSetTodoDone(t *testing.T) {
	setup()
	req, _ := http.NewRequest("PATCH", "/todo/0", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	req, _ = http.NewRequest("GET", "/todo/0", nil)
	router.ServeHTTP(rec, req)
	expected := "{\"Id\":0,\"Description\":\"Write an API\",\"DueDate\":\"07.05.2021\",\"Done\":true,\"Priority\":10}\n"
	assert.Equal(t, expected, rec.Body.String())
}
