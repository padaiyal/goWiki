package main

import "fmt"

func squareNumbers(numbers []int) []int {
	for index, number := range numbers {
		numbers[index] = number * number
	}
	return numbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, -6}
	// Pass by reference
	fmt.Printf("Numbers: %v\n", numbers)
	squaredNumbers := squareNumbers(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Squared numbers: %v\n", squaredNumbers)
}
