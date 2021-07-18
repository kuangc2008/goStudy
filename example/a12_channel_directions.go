package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ping(pings chan<- string, s string) {
	pings <- s
}
