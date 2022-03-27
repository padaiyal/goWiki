package main

import (
	"fmt"
	"strings"
)

func sayGoodBye(name string) {
	fmt.Printf("Goodbye %s!", name)
}

func main() {
	// Deferred Function Call Argument Evaluation.
	names := []string{"Albus", "Percival", "Wulfric", "Brian", "Dumbledore"}
	defer sayGoodBye(strings.Join(names, " "))
	names[0] = "Megabus"
}
