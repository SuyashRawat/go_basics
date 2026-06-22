package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	fmt.Println("Starting to print numbers...")

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 200)
	}

	fmt.Println("Finished printing numbers.")

}

func main() {

	go printNumbers()

	time.Sleep(time.Second * 2)

	for i := 0; i < 5; i++ {

		go func() {
			fmt.Println("without capturing", i)
		}()

	}

	for i := 0; i < 5; i++ {

		go func(i int) {
			fmt.Println("with capturing", i)
		}(i)

	}

	time.Sleep(time.Second * 2)

}
