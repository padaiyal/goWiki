package main

import "fmt"

func main() {
	// Defer order - Follows LIFO like a stack.
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
}
