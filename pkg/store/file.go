package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/hunter32292/tasky/pkg/model"
)

var (
	taskDBLocation = ".tasky.db"
	// Tasks - In memory collection of tasks
	Tasks = []*model.Task{}
)

// Save - in memory tasks
func Save() {
	file, err := os.Create(createInHomeDir(taskDBLocation))
	if err != nil {
		log.Fatal("Failed to access tasky.db: ", err)
	}
	dat, err := json.Marshal(Tasks)
	if err != nil {
		log.Fatal("Could not read in-memory tasks: ", err)
	}

	bytes, err := file.Write(dat)
	if err != nil {
		log.Fatal("Failed to write to tasky.db: ", err)
	}

	log.Println("Save tasks, wrote bytes -", bytes)
}

// Load - saved tasks from file
func Load() {
	// Check if file exists
	_, err := os.Stat(createInHomeDir(taskDBLocation))
	if os.IsNotExist(err) {
		log.Println("tasky.db doesn't exist, creating")
		_, _ = os.Create(createInHomeDir(taskDBLocation))
		return
	}
	// Load file
	file, err := os.Open(createInHomeDir(taskDBLocation))
	if err != nil {
		log.Fatal("Failed to access tasky.db: ", err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read tasky.db: ", err)
	}
	if len(bytes) <= 0 {
		return
	}
	err = json.Unmarshal(bytes, &Tasks)
	if err != nil {
		log.Fatal("Failed to create in-memory tasks from tasky.db: ", err)
	}
}

func createInHomeDir(filename string) string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	filepath := fmt.Sprintf("%s/%s", user.HomeDir, filename)
	return filepath
}
