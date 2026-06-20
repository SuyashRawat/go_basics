package main

import "fmt"

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

// variadic function
func sum(nums ...int) int {

	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

func main() {
	q, r := divide(10, 3)

	fmt.Println(q)
	fmt.Println(r)

	s, p := calculate(2, 3)

	fmt.Println(s, p)

	fmt.Println(sum(1, 2, 3, 4))

	func() {
		fmt.Println("Anonymous Function")
	}()
}
