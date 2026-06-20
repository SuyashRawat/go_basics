package main

import "fmt"

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
)

const Pi = 3.14159

func main() {
	var name string = "Suyash"
	var age int = 22

	fmt.Println(name)
	fmt.Println(age)
	var a int = 10
	var b float64 = 10.5
	var c bool = true
	var d string = "golang"

	fmt.Println(a, b, c, d)

	var userName string

	fmt.Print("Enter name: ")
	fmt.Scan(&userName)

	fmt.Println("Hello", userName)

}
