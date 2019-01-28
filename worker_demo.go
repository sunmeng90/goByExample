package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker:", id, "starting on job:", job)
		time.Sleep(500 * time.Millisecond)
		results <- job
		fmt.Println("Worker:", id, "finished for job:", job)
	}
}

//http://www.golangbootcamp.com/book/concurrency
func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	// fatal error: all goroutines are asleep - deadlock!
	// That’s because we overfilled the buffer without letting the code a chance to read/remove a value from the channel.
	// Without the buffer, the channel is overfilled when keeping send value to it without consuming
	// jobs := make(chan int)
	// results := make(chan int)
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 0; a < 5; a++ {
		result := <-results
		fmt.Println("Receive results: ", result)
	}
	fmt.Scanln()

}
