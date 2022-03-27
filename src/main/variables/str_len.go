package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// len() returns the number of bytes, not the length of the string
	text := "☺"
	fmt.Println(len(text))
	fmt.Println(utf8.RuneCountInString(text))
}
