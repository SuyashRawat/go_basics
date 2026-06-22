package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Sum(10, 20))
	fmt.Println(Sum(10.5, 20.3))
}
