package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		printUsage()
		return
	}
	command := os.Args[1]

	switch command {
	case "help":
		printUsage()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description.")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks(getTasks())
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide task ID and new status.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		updateTask(id, os.Args[3])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID to delete.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		deleteTask(id)
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID to mark as in-progress.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		markTask(id, "in-progress")
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID to mark as done.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		markTask(id, "done")
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
func updateTask(id int, description string) {
	tasks := getTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Task updated successfully.")
			return
		}
	}
	fmt.Println("Error: Task with ID", id, "not found.")
}
func deleteTask(id int) {
	tasks := getTasks()
	var newTasks []Task
	found := false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}
	if !found {
		fmt.Println("Error: Task with ID", id, "not found.")
		return
	}
	saveTasks(newTasks)
	fmt.Println("Task deleted successfully.")
}
func markTask(id int, status string) {
	tasks := getTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Task marked as", status, "successfully.")
			return
		}
	}
	fmt.Println("Error: Task with ID", id, "not found.")
}
func printUsage() {
	fmt.Println("Usage: task-cli <command> [<args>]")
	fmt.Println("Commands:")
	fmt.Println("  add <description>         Add a new task")
	fmt.Println("  list                      List all tasks")
	fmt.Println("  update <id> <description> Update task description")
	fmt.Println("  delete <id>               Delete a task")
	fmt.Println("  mark-in-progress <id>     Mark a task as in-progress")
	fmt.Println("  mark-done <id>            Mark a task as done")
	fmt.Println("  help                      Show this help message")
}
