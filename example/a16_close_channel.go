package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			fmt.Println("for 1")
			// s. In this special 2-value form of receive,
			//the more value will be false if jobs has been closed and all values in the channel have already been received.
			j, more := <-jobs
			fmt.Println("for 2")
			if more {
				fmt.Println("receive", j)
			} else {
				fmt.Println("all job")
				done <- true
				return
			}

		}
	}()

	time.Sleep(time.Second)

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)
	fmt.Println("send all jobs")

	for elm := range jobs {
		fmt.Println("elm:", elm)
	}

	<-done
}
