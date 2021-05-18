package model

type PostTodo struct {
	Description string
	DueDate     string
	Priority    int
}

type Todo struct {
	Id          int
	Description string
	DueDate     string
	Done        bool
	Priority    int
}

type Todos []Todo

func (todos Todos) Len() int {
	return len(todos)
}
func (todos Todos) Less(i, j int) bool {
	return todos[i].Id < todos[j].Id
}
func (todos Todos) Swap(i, j int) {
	todos[i], todos[j] = todos[j], todos[i]
}
