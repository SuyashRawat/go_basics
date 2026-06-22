package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialized")
}

func main() {

	for i := 0; i < 5; i++ {

		go func() {
			once.Do(initialize)
		}()
	}

	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			atomic.AddInt64(
				&counter,
				1,
			)

		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
