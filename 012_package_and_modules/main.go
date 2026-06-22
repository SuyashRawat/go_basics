package main

import (
	"fmt"
	"go_basics/012_package_and_modules/calculator"
)

func main() {
	a := 10
	b := 5

	sum := calculator.Add(a, b)
	product := calculator.Multiply(a, b)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Product: %d\n", product)
}
