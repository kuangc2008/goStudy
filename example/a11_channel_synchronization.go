package main

import (
	"fmt"
	"time"
)

func main() {
	//done := make(chan bool, 1)
	//done := make(chan bool)
	done := make(chan bool, 2)
	go work(done)
	//done <- true
	// Block until we receive a notification from the worker on the channel.
	fmt.Println(<-done)
	fmt.Println("bbb")
}

func work(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second * 2)
	fmt.Println("done")

	// The done channel will be used to notify another goroutine that this function’s work is done.
	done <- true

}
