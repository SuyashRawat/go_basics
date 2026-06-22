package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan any)

	go func() {
		ch <- 42
	}()

	val := <-ch

	fmt.Println(val)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch2 <- "done"
	}()

	select {

	case msg := <-ch2:
		fmt.Println("message is ", msg)

	case <-time.After(5400 * time.Millisecond):
		fmt.Println("timeout")
	}

}
