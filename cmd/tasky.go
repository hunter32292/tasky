package cmd

import (
	"fmt"
	"os"

	"github.com/hunter32292/tasky/pkg/store"
	"github.com/hunter32292/tasky/pkg/task"
)

const (
	// application name
	applicationName    = "tasky"
	applicationHelpMsg = `
	Tasky is a simple and extensible task manager for daily use in organizing your day to day tasks.
	Example usage:
	tasky create
		title : Pick up groceries
		content : # Grocery List
		1. Eggs
		1. Fruit cups
		1. Milk
		due date: 5d
	Task created:
	Pick Up Groceries
	
	Grocery List
	- Eggs
	- Fruit cups
	- Milk

	Due - Tue, 09 Mar 2021 15:56:39 -0800
	`
)

// RunTasky - Run the task generator command
func RunTasky() {
	// Load task file
	store.Load()

	if len(os.Args) <= 1 {
		fmt.Println(applicationHelpMsg)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "create":
		task.CreateTask()
	case "list":
		task.ListTask()
	case "clear":
		task.ClearTask()
	default:
		fmt.Println(applicationHelpMsg)
	}

	// Save at end of process
	store.Save()
}
