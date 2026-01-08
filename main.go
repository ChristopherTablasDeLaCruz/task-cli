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
		addTask(os.Args[2])
	case "list":
		listTasks(getTasks())
	default:
		fmt.Println("Error: Unknown command:", command)
	}
}

func addTask(description string) {
	tasks := getTasks()
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	newTask := Task{
		ID:          maxID + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Println("Task added successfully.")
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
func getTasks() []Task {
	var tasks []Task
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return []Task{}
	}
	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Error unmarshaling tasks:", err)
		return []Task{}
	}
	return tasks
}
func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Printf("%-5s %-25s %-5s\n", "ID", "Description", "Status")
	fmt.Println("---------------------------------------")
	for _, task := range tasks {
		fmt.Printf("%-5d %-25s %-5s\n",
			task.ID, task.Description, task.Status)
	}
}
