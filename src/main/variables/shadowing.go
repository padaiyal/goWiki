package main

import "fmt"

func arePositiveNumbers(numbers []int) bool {
	result := true
	for _, number := range numbers {
		if number < 0 {
			result := false
			fmt.Printf("result: %t\n", result)
			break
		}
	}
	return result
}

func main() {
	// Variable shadowing
	numbers := []int{1, 2, 3, 4, 5, -6}
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Actual result: %t\n", arePositiveNumbers(numbers))
}
