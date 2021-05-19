package database

import (
	"homework-SoerenDev-391298709521/model"
	"sort"
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
