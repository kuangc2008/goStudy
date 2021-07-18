package main

import "fmt"

func main() {
	message := make(chan string)
	signal := make(chan bool)

	select {
	// If a value is available on messages then select will take the <-messages case with that value.
	//If not it will immediately take the default case.
	case msg := <-message:
		fmt.Println("receiverd message", msg)
	default:
		fmt.Println("no message recivied")
	}

	msg := "hi"
	select {
	// Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
	//Therefore the default case is selected.
	case message <- msg:
		fmt.Println("send messsage", msg)
	default:
		fmt.Println("no send")
	}

	select {
	case msg := <-message:
		fmt.Println("rec3", msg)
	case sig := <-signal:
		fmt.Println("signal", sig)
	default:
		fmt.Println("no activity")
	}
}
