package main

import "fmt"

func main() {
	var nums [5]int

	fmt.Println(nums)

	nums2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(nums2)

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)

	arr1 := [3]int{1, 2, 3}
	arr2 := arr1

	arr2[0] = 100

	fmt.Println(arr1) //1 2 3
	fmt.Println(arr2) //100 2 3
}
