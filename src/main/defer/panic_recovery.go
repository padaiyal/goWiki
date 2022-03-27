package main

func main() {
	// Recovering from a panic.

	//recover()
	//panic("OMG")
	//fmt.Println("After panic ....") // Will not reach here, so can't recover here.

	// The panic recovery should be in a deferred function.
	//defer func() {
	//	fmt.Println("Recovered: ", recover())
	//}()
	//panic("OMG")
}
