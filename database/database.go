package database

import (
	"errors"
	"sync"
	"todolist/model"
)

type TodoRepository struct {
	data model.Todos
	sync.RWMutex
}

func NewTodoRepository() TodoRepository {
	return TodoRepository{}
}

func (repo *TodoRepository) AddTodo(todo model.PostTodo) {
	id := findNewId(repo.data)
	repo.data = append(repo.data, model.Todo{Id: id, DueDate: todo.DueDate, Description: todo.Description, Priority: todo.Priority})
}

func (repo *TodoRepository) GetAllTodos() []model.Todo {
	return repo.data
}

func (repo *TodoRepository) GetTodoById(id int) (model.Todo, error) {
	for i := range repo.data {
		if repo.data[i].Id == id {
			return repo.data[i], nil
		}
	}
	return model.Todo{}, errors.New("id not found")
}

func (repo *TodoRepository) TodoIsDone(id int) error {
	for i := range repo.data {
		if repo.data[i].Id == id {
			repo.data[i].Done = true
			return nil
		}
	}
	return errors.New("id not found")
}

func (repo *TodoRepository) DeleteTodo(id int) error {
	for i := range repo.data {
		if repo.data[i].Id == id {
			repo.data[i] = repo.data[len(repo.data)-1]
			repo.data = repo.data[:len(repo.data)-1]
			return nil
		}

	}
	return errors.New("id not found")
}
