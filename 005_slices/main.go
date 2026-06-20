package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(nums)

	nums = append(nums, 5)
	fmt.Println(nums)

	s := make([]int, 5)

	fmt.Println(s)
	fmt.Println("len =", len(s))
	fmt.Println("cap =", cap(s))

	src := []int{1, 2, 3}

	dst := make([]int, len(src))

	copy(dst, src)

	dst[0] = 100

	fmt.Println(src)
	fmt.Println(dst)
}
