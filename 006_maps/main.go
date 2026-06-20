package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["apple"] = 10

	fmt.Println(m)

	m["apple"] = 100

	fmt.Println(m["apple"])

	m["banana"] = 20

	for key, value := range m {
		//Order is not guaranteed
		fmt.Printf("%s: %d\n", key, value)
	}

	delete(m, "banana")

	val, ok := m["apple"]

	fmt.Println(val, ok)

	val, ok = m["orange"]

	fmt.Println(val, ok)
}
