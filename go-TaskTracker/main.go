package main

import (
	"fmt"
	"go-TaskTracker/internal/controller"
	"go-TaskTracker/pkg/utils"
	"os"
	"strconv"
)

type Tasks struct {
	status string
	name   string
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No command passed !")
		return
	}

	switch args[0] {
	case "add":
		if len(args) > 2 {
			fmt.Println("Usage: task-cli add 'Task description'")
			return
		}
		description := args[1]
		task := controller.AddTask(description)
		fmt.Printf("Task added successfully (ID: %v)\n", task.ID)

	case "update":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli update <id> 'new description'")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		newDescription := args[2]
		controller.UpdateTask(id, newDescription)
		fmt.Printf("Task updated successfully (ID: %v)\n", id)
		return

	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		utils.CheckNilError(err)

		_, msg := controller.DeleteTask(id)

		if msg != "success" {
			fmt.Printf("Task doesn't exist")
			return
		}
		fmt.Println("task deleted successfully!")
		return

	case "mark-in-progress":
		if len(args) > 2 {
			fmt.Println("Usage: task-cli mark in progress <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		controller.MarkTask(id, "inprogress")
		fmt.Printf("Task marked as in progress (ID: %v)\n", id)

	case "mark-done":
		if len(args) > 2 {
			fmt.Println("Usage: task-cli mark done <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		controller.MarkTask(id, "completed")
		fmt.Printf("Task marked as completed (ID: %v)\n", id)

	default:
		fmt.Print("Unkown command :(")
	}

}
