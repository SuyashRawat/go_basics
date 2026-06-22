package main

import (
	"fmt"
	"sync"
)

func worker(
	id int,
	jobs <-chan int,
	results chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf(
			"Worker %d processing %d\n",
			id,
			job,
		)

		results <- job * job
	}
}

func main() {

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go worker(
			i,
			jobs,
			results,
			&wg,
		)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	close(results)

	for result := range results {
		fmt.Println("result is ", result)
	}
}
