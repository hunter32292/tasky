package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/hunter32292/tasky/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.RunTasky()
	os.Exit(0)
}
