package database

import (
	"sort"
	"todolist/model"
)

func findNewId(todos model.Todos) int {
	sort.Sort(todos)
	for i := range todos {
		if i != todos[i].Id {
			return i
		}
	}
	return todos.Len()
}
