package main

import (
	"fmt"
	"time"
)

func performTask(done chan bool) {
	time.Sleep(4 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)

	go performTask(done)

	select {
	case <-done:
		fmt.Println("Task completed successfully")
	case <-time.After(3 * time.Second):
		fmt.Println("Task timed out")
	}

	fmt.Println("Program exiting")
}
