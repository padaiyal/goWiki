package main

import (
	"fmt"
	"time"
)

func runTask(taskId int) {
	fmt.Printf("[%v] is running\n", taskId)
	time.Sleep(1 * time.Second)
	fmt.Printf("[%v] is done\n", taskId)
}

func main() {
	// App exits with active go routines.
	taskCount := 4

	for i := 0; i < taskCount; i++ {
		go runTask(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("All tasks completed.")

	// Use a wait group to prevent premature termination of go routines.
	//var waitGroup sync.WaitGroup
	//taskCount := 4
	//
	//for i := 0; i < taskCount; i++ {
	//	waitGroup.Add(1)
	//	go func() {
	//		defer waitGroup.Done()
	//		runTask(i)
	//	}()
	//
	//	waitGroup.Wait()
	//}
	fmt.Println("All tasks completed.")
}
