# Todoist Go API Wrapper

![Go](https://img.shields.io/badge/Go-1.23-blue)
![License](https://img.shields.io/badge/license-MIT-green)

A simple Go client for interacting with the [Todoist REST API](https://developer.todoist.com/rest/v2/). This wrapper provides an easy-to-use interface for managing projects, tasks, labels, sections, and comments in Todoist.

## Features

- **Projects**: Create, update, retrieve, and delete projects.
- **Tasks**: Create, update, retrieve, close, and delete tasks.
- **Sections**: Manage sections within projects.
- **Labels**: Handle personal and shared labels.
- **Comments**: Add, update, and delete comments on tasks and projects.

## Installation

You can install the package using `go get`:

```bash
go get github.com/felixschmelzer/todoist-go
```

Then, import it in your Go project:
```go
import "github.com/felixschmelzer/todoist-go"
```

## Usage

### Getting Your Todoist API Token

Before you can use this Go API wrapper, you'll need to obtain your Todoist API token. The token allows you to authenticate your requests and interact with the Todoist API.
How to get your API Token is described [here](https://todoist.com/de/help/articles/find-your-api-token).

Once you have the token, you can use it to initialize the TodoistClient in your Go code and access Todoist Resources:

```go
package main

import (
	"fmt"
	"github.com/felixschmelzer/todoist-go"
)

func main() {
	// Create Client with your token
	const token = "your_todoist_api_token"
	client := todoist.NewTodoistClient(token)
	
	// Get all Projects
	projects, err := client.GetProjects()
	if err != nil {
		fmt.Println("Error fetching projects:", err)
		return
	}

	for _, p := range projects {
		fmt.Printf("Project: %s (ID: %s)\n", p.Name, p.ID)
	}

	// Create a New Task
	taskParams := todoist.TaskParams{
		Content:   "Buy groceries",
		DueString: "tomorrow at 12:00",
		Priority:  4,
	}

	task, err := client.CreateTask(taskParams)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return
	}

	fmt.Printf("Created task: %s (ID: %s)\n", task.Content, task.ID)

	// Update a Task
	updateParams := todoist.TaskParams{
		Content: "Buy groceries and milk",
	}

	updatedTask, err := client.UpdateTask(task.ID, updateParams)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return
	}

	fmt.Printf("Updated task: %s (ID: %s)\n", updatedTask.Content, updatedTask.ID)

	// Get all Tasks
	tasks, err := client.GetTasks("", "", "")
	if err != nil {
		fmt.Println("Error getting Tasks:", err)
		return
	}

	for _, t := range tasks {
		fmt.Printf("Task: %s (ID: %s)\n", t.Description, t.ID)
	}
}
```


## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to open an issue or submit a pull request.

## Acknowledgments

This package is inspired by the official Todoist REST API and aims to make interacting with Todoist easier for Go developers.

