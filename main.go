package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [<args>]")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description.")
			return
		}
		description := os.Args[2]
		var tasks []Task
		data, err := os.ReadFile("tasks.json")
		if err == nil {
			json.Unmarshal(data, &tasks)
		}
		newTask := Task{
			ID:          len(tasks) + 1,
			Description: description,
			Status:      "pending",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, newTask)
		saveTasks(tasks)
		fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	case "list":
		fmt.Println("Listing all tasks...")
	default:
		fmt.Println("Error: Unknown command:", command)
	}
}

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
}
