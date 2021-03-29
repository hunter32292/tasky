package task

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hunter32292/tasky/pkg/model"
	"github.com/hunter32292/tasky/pkg/store"
)

// ClearTask - Remove tasks from the task list
func ClearTask() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please provide a Task ID to remove:")
	taskID, _ := reader.ReadString('\n')
	taskID = strings.ReplaceAll(taskID, "\n", "")
	if taskID == "" {
		fmt.Println("This will clear all tasks [Enter] to continue...")
		fmt.Println("or (S)kip...")
		taskID, _ := reader.ReadString('\n')
		taskID = strings.ReplaceAll(taskID, "\n", "")
		if taskID == "S" {
			return
		}
		clearTask(-1)
		return
	}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		log.Fatal(err)
	}
	clearTask(id)
}

func clearTask(id int) []*model.Task {
	if id == -1 {
		store.Tasks = []*model.Task{}
	} else {
		store.Tasks = append(store.Tasks[:id], store.Tasks[id+1:]...)
	}
	return store.Tasks
}

// ListTask - Print of task list
func ListTask() {
	Tasks := listTask()
	for index, task := range Tasks {
		fmt.Printf("%d)\t Title:%s \tDue Date:%s\n", index, task.Title, task.DueDate.Format(time.RFC1123Z))
	}
}

func listTask() []*model.Task {
	return store.Tasks
}

// CreateTask - Handle create of a task
func CreateTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Title :")
	title, _ := reader.ReadString('\n')
	title = strings.ReplaceAll(title, "\n", "")
	fmt.Printf("Content :")
	content, _ := reader.ReadString('\n')
	fmt.Printf("Due Date :")
	due := getDueDate()
	// Create and add a task to the task list
	store.Tasks = append(store.Tasks, createTask(title, content, due))
}

func createTask(title string, content string, dueDate time.Time) *model.Task {
	return &model.Task{Title: title, Content: content, DueDate: dueDate}
}

func getDueDate() time.Time {
	currentTime := time.Now()
	reader := bufio.NewReader(os.Stdin)
	duedate, _ := reader.ReadString('\n')
	duedate = strings.ReplaceAll(duedate, "\n", "")
	arr := strings.Split(duedate, "")
	lastValue, x := arr[len(arr)-1], arr[:len(arr)-1]
	value, err := strconv.Atoi(strings.Join(x, ""))
	if err != nil {
		log.Fatal(err)
	}
	multiFactor := int64(value)
	switch lastValue {
	case "h":
		addTime := int64(time.Hour) * multiFactor
		currentTime = currentTime.Add(time.Duration(addTime))
	case "d":
		addTime := (int64(time.Hour) * 24) * multiFactor
		currentTime = currentTime.Add(time.Duration(addTime))
	case "w":

		addTime := ((int64(time.Hour) * 24) * 7) * multiFactor
		currentTime = currentTime.Add(time.Duration(addTime))
	default:
		fmt.Println("Did not understand time values provided")
		fmt.Printf("Please use(where X is number):\n hours = Xh\n days = Xd\n weeks = Xw\n")
	}
	return currentTime
}
