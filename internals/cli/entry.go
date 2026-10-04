package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const todosFile = "assets/todos.json"

func SrartCli(){
		if len(os.Args) < 2 {
		fmt.Println("usage: todo <add|list|toggle> [args]")
		return
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: todo add <text>")
			return
		}
		body := strings.Join(os.Args[2:], " ")
		if err := AddTodo(Todo{Body: body}, todosFile); err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Printf("added %q\n", body)

	case "list":
		todos, err := ReadTodos(todosFile)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		if len(todos) == 0 {
			fmt.Println("no todos yet")
			return
		}
		for _, t := range todos {
			check := " "
			if t.Done {
				check = "x"
			}
			fmt.Printf("%d. [%s] %s\n", t.ID, check, t.Body)
		}

	case "toggle":
		if len(os.Args) < 3 {
			fmt.Println("usage: todo toggle <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id must be a number")
			return
		}
		todo, err := ToggleTodo(id, todosFile)
		if errors.Is(err, ErrTodoNotFound) {
			fmt.Printf("no todo with id %d\n", id)
			return
		}
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		status := "not done"
		if todo.Done {
			status = "done"
		}
		fmt.Printf("%q marked as %s\n", todo.Body, status)

	default:
		fmt.Printf("unknown command %q\n", os.Args[1])
		fmt.Println("usage: todo <add|list|toggle> [args]")
	}
}
