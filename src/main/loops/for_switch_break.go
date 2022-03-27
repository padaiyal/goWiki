package loops

import "fmt"

func main() {
	// Breaking Out of "for switch" and "for select" Code Blocks
	// Sum of all numbers. If any of them is the max value of the uint type,
	// just set the sum to the max value and break.
	maxUint := ^uint(0)
	numbers := []uint{10, 11, 12, maxUint, 1, 2, 3, 4}
	var sum uint = 0
	for _, number := range numbers {
		switch number {
		case maxUint:
			sum = maxUint
			break
		default:
			sum += number
		}
	}
	fmt.Println(sum)

	// Add a label enclosing the loop and specify it to break out of the loop.
	//	sum = 0
	//loop:
	//	for _, number := range numbers {
	//		switch number {
	//		case maxUint:
	//			sum = maxUint
	//			break loop
	//		default:
	//			sum += number
	//		}
	//	}
	//	fmt.Println(sum)
}
