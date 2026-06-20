package main

import "fmt"

type Counter struct {
	Value int
}

// Value Receiver
func (c Counter) IncrementWrong() {
	c.Value++
}

// Pointer Receiver
func (c *Counter) IncrementCorrect() {
	c.Value++
}

func (c Counter) Print() {
	fmt.Println(c.Value)
}

func main() {

	c := Counter{}

	c.IncrementWrong()
	c.Print()

	c.IncrementCorrect()
	c.Print()
}
