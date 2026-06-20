package main

import "fmt"

func greet() {
	fmt.Println("Hello")
}

func add(a int, b int) int {
	return a + b
}

func main() {
	greet()
	fmt.Println(add(2, 3))
}
