package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		// sends and receives block until both the sender and receiver are ready.
		messages <- "ping"
	}()

	msg := <-messages

	fmt.Println(msg)

}
