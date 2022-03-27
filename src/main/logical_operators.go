package main

import "fmt"

func main() {
	// NOT operator VS XOR operator
	fmt.Println(^1)    // ^0|01 -> 1|10
	fmt.Println(1 ^ 2) // 01 ^ 10 -> 11
}
