package main

import (
	"fmt"
)

func main() {
	// "Hidden" Data in Slices
	names := []string{"Alice", "Bob", "Ajay", "Cheng"}
	newNames := names[:2]
	fmt.Println(len(names), cap(names), &names[0])
	fmt.Println(len(newNames), cap(newNames), &newNames[0])
}
