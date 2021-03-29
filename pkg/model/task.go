package model

import "time"

// Task - Task object to be finished last on
type Task struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	DueDate  time.Time `json:"duedate"`
	SubTasks []*Task   `json:"subtasks"`
}
