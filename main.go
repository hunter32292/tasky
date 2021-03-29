package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hunter32292/tasky/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("Starting ...")
	cmd.RunTasky()
	os.Exit(0)
}
