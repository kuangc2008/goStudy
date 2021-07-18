package main

import "fmt"

func main() {

	// Buffered channels accept a limited number of values without a corresponding receiver for those values.
	message := make(chan string, 2)

	// Note that the channel is buffered, so the send in the goroutine is nonblocking.
	message <- "buffered"
	//message <- "channel"
	message <- "channel2"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
