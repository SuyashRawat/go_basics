package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Address struct {
	City string
}

type Employee struct {
	User
	Address
	Salary int
}

func main() {

	// -------------------------
	// Struct Creation
	// -------------------------

	u1 := User{
		Name: "Suyash",
		Age:  22,
	}

	fmt.Println(u1)

	// -------------------------
	// Positional initialization
	// -------------------------

	u2 := User{"Rahul", 25}

	fmt.Println(u2)

	// -------------------------
	// Anonymous Struct
	// -------------------------

	config := struct {
		Port  int
		Debug bool
	}{
		Port:  8080,
		Debug: true,
	}

	fmt.Println(config)

	// -------------------------
	// Embedded Struct
	// -------------------------

	e := Employee{
		User: User{
			Name: "John",
			Age:  30,
		},
		Address: Address{
			City: "Delhi",
		},
		Salary: 100000,
	}

	fmt.Println(e.Name)
	fmt.Println(e.City)

	// -------------------------
	// Struct Comparison
	// -------------------------

	a := User{"A", 20}
	b := User{"A", 20}

	fmt.Println(a == b)
}
