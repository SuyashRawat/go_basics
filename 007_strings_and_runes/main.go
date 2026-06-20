package main

import (
	"fmt"
	"strings"
)

func main() {

	// -------------------------
	// String Basics
	// -------------------------

	s := "hello"

	fmt.Println(len(s)) // bytes

	// -------------------------
	// Strings are immutable
	// -------------------------

	// s[0] = 'H' // compile error

	// -------------------------
	// String vs Rune
	// -------------------------

	word := "你好"

	fmt.Println("bytes =", len(word))
	fmt.Println("runes =", len([]rune(word)))

	// -------------------------
	// Iterate over bytes
	// -------------------------

	str := "नमस्ते"

	fmt.Println("\nByte iteration")

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}

	// -------------------------
	// Iterate over runes
	// -------------------------

	fmt.Println("\n\nRune iteration")

	for idx, r := range str {
		fmt.Printf("%d -> %c\n", idx, r)
	}

	// -------------------------
	// String Builder
	// -------------------------

	var builder strings.Builder

	builder.WriteString("Go")
	builder.WriteString("Lang")

	fmt.Println("\nBuilder:", builder.String())

	// -------------------------
	// Byte Slice Conversion
	// -------------------------

	data := []byte("hello")

	data[0] = 'H'

	fmt.Println(string(data))
}
