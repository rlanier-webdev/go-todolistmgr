package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a single to-do item with a description and a completion status
type Task struct {
	Description string
	Status      bool
}

func main() {
	// Initialize an empty slice of tasks, where each task is of the Task type
	tasks := []Task{}

	// Create a new buffered reader to read input from the standard input
	scanner := bufio.NewReader(os.Stdin)

	var choice int

	// Main loop for user interaction
	for {
		fmt.Println("1. Add a new task")
		fmt.Println("2. View all tasks")
		fmt.Println("3. Mark a task as completed")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Add task
			fmt.Print("Enter task description: ")
			// Reads a line of text from the standard input up to and including the newline character
			description, err := scanner.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			// Trim whitespace from the description
			description = strings.TrimSpace(description)
			task := Task{Description: description, Status: false}
			addTask(&tasks, task)
		case 2:
			viewTasks(tasks)
		case 3:
			markTaskAsCompleted(&tasks)
		case 4:
			deleteTask(&tasks)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again")
		}
	}
}

// Adds a new task to the list of tasks
func addTask(tasks *[]Task, task Task) {
	if task.Description == "" {
		fmt.Println("Task description cannot be empty.")
		return
	}
	*tasks = append(*tasks, task)
	fmt.Println("Task added successfully.")
}

// Displays all tasks with their completion status
func viewTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to display.")
		return
	}
	fmt.Println("Task(s):")
	for i, task := range tasks {
		status := "Incomplete"
		if task.Status {
			status = "Complete"
		}
		fmt.Printf("%d. %s - %s\n", i+1, task.Description, status)
	}
	fmt.Println("---------------------------")
}

// Marks a specific task as completed based on user input
func markTaskAsCompleted(tasks *[]Task) {
	var index int
	fmt.Print("Enter task number to mark as completed: ")
	fmt.Scanln(&index)
	if index > 0 && index <= len(*tasks) {
		(*tasks)[index-1].Status = true
		fmt.Println("Task marked as completed.")
	} else {
		fmt.Println("Invalid task number.")
	}
}

// Deletes a specific task based on user input
func deleteTask(tasks *[]Task) {
	var index int
	fmt.Print("Enter task number to delete: ")
	fmt.Scanln(&index)
	if index > 0 && index <= len(*tasks) {
		*tasks = append((*tasks)[:index-1], (*tasks)[index:]...)
		fmt.Println("Task deleted.")
	} else {
		fmt.Println("Invalid task number.")
	}
}
