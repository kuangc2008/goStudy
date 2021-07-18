package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a < numJobs; a++ {
		fmt.Println("heihei", a)
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("len:", len(jobs))
	for j := range jobs {
		fmt.Println("workder", id, "statrt job ", j)
		time.Sleep(time.Second)
		fmt.Println("workder", id, "finish job ", j)
		results <- j * 2

	}
}
