package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"` // todo, in-progress, done
}

func loadTasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	return os.WriteFile("tasks.json", data, 0644)
}

func addTask(desc string) {
	tasks, _ := loadTasks()

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task := Task{
		ID:          id,
		Description: desc,
		Status:      "todo",
	}

	tasks = append(tasks, task)
	saveTasks(tasks)

	fmt.Println("Task added!")
}

func updateStatus(id int, status string) {
	tasks, _ := loadTasks()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
		}
	}
	fmt.Println("Status Updated")

	saveTasks(tasks)
}

func deleteStatus(id int) {
	tasks, _ := loadTasks()

	var newTasks []Task

	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		}
	}

	saveTasks(newTasks)
	fmt.Println("Task Deleted")
}

func listTasks(filter string) {
	tasks, _ := loadTasks()

	for _, t := range tasks {
		if filter == "" || t.Status == filter {
			fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
		}
	}
}
func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No command")
		return
	}

	switch args[1] {

	case "add":
		addTask(args[2])

	case "list":
		if len(args) == 3 {
			listTasks(args[2])
		} else {
			listTasks("")
		}

	case "done":
		id, _ := strconv.Atoi(args[2])
		updateStatus(id, "done")

	case "progress":
		id, _ := strconv.Atoi(args[2])
		updateStatus(id, "in-progress")

	case "delete":
		id, _ := strconv.Atoi(args[2])
		deleteStatus(id)

	default:
		fmt.Println("Unknown command")
	}
}
