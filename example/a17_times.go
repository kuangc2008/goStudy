package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(2 * time.Second)
	<-time1.C
	fmt.Println("timer 1 fired")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("time2 fireed")
	}()

	// stop
	stop := time2.Stop()
	if stop {
		fmt.Println("stop")
	}

	time.Sleep(2 * time.Second)

}
