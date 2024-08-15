# Task Management CLI Program

A simple command-line interface (CLI) program written in Go that allows users to manage a list of tasks. Users can add new tasks, view existing tasks, mark tasks as completed, and delete tasks.

## Features

- **Add New Task:** Allows the user to add a new task by providing a description.
- **View All Tasks:** Displays all tasks with their completion status.
- **Mark Task as Completed:** Lets the user mark a specific task as completed.
- **Delete Task:** Enables the user to delete a specific task from the list.
- **User-friendly Interface:** Provides a menu-driven interface for ease of use.

## Getting Started

### Prerequisites

- Go installed on your machine. You can download it from [Go's official website](https://golang.org/dl/).

### Installation

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/rlanier-webdev/go-todolistmgr.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go-todolistmgr
   ```
3. Build the program:
   ```bash
   go build
   ```

### Usage

1. Run the program:
   ```bash
   ./go-todolistmgr
   ```
2. Follow the on-screen menu to manage your tasks.

### Example

```plaintext
1. Add a new task
2. View all tasks
3. Mark a task as completed
4. Delete a task
5. Exit
Enter your choice: 1
Enter task description: Buy groceries
Task added successfully.
```
