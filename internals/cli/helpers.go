package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"slices"
)

type Todo struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
	Done bool   `json:"done"`
}

func ReadTodos(path string) ([]Todo, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return []Todo{}, nil
	}
	if err != nil {
		return nil, err
	}

	var todos []Todo

	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func AddTodo(todo Todo, path string) error {
	todos, err := ReadTodos(path)
	if err != nil {
		return err
	}

	nextId := 1

	for _, t := range todos {
		if t.ID >= nextId {
			nextId = t.ID + 1
		}
	}
	todo.ID = nextId
	todos = append(todos, todo)

	if err := saveTodos(todos, path); err != nil {
		return err
	}
	return nil

}

var ErrTodoNotFound = errors.New("todo not found")

func ToggleTodo(id int, path string) (Todo, error) {
	todos, err := ReadTodos(path)
	if err != nil {
		return Todo{}, err
	}
	i := slices.IndexFunc(todos, func(t Todo) bool { return t.ID == id })
	if i == -1 {
		return Todo{}, fmt.Errorf("%w id %d", ErrTodoNotFound, id)
	}
	todos[i].Done = !todos[i].Done

	if err := saveTodos(todos, path); err != nil {
		return Todo{}, err
	}
	return todos[i], nil

}

func saveTodos(todos []Todo, path string) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
