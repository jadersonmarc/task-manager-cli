package main

import (
	"fmt"
	"os"

	"github.com/jadersonmarc/task-manager-cli/service"
	"github.com/jadersonmarc/task-manager-cli/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("uso: task-cli <comando>")
		return
	}

	cmd := os.Args[1]

	switch cmd {

	case "add":
		err := service.Add(os.Args[2])
		handleErr(err)

	case "done":
		err := service.UpdateStatus(os.Args[2], task.StatusDone)
		handleErr(err)

	case "progress":
		err := service.UpdateStatus(os.Args[2], task.StatusInProgress)
		handleErr(err)

	case "delete":
		err := service.Delete(os.Args[2])
		handleErr(err)

	case "list":
		print(service.List(func(t task.Task) bool { return true }))

	case "list:done":
		print(service.List(func(t task.Task) bool {
			return t.Status == task.StatusDone
		}))

	case "list:todo":
		print(service.List(func(t task.Task) bool {
			return t.Status == task.StatusTodo
		}))

	case "list:progress":
		print(service.List(func(t task.Task) bool {
			return t.Status == task.StatusInProgress
		}))
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Sucesso!")
}

func print(tasks []task.Task, err error) {
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	for _, t := range tasks {
		fmt.Printf("[%s] %d - %s\n", t.Status, t.ID, t.Description)
	}
}
