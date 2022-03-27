package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Albus", "Percival", "Wulfric", "Brian", "Dumbledore"}
	// Iteration variable closure when using subroutines.
	for index, name := range names {
		go func() {
			fmt.Println(index, name)
		}()
	}
	time.Sleep(2 * time.Second)
}
