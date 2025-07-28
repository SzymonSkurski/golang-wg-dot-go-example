package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// run function in goroutine - old way
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Running in goroutine - old way")
	}()

	// run function in goroutine - new way
	wg.Go(func() {
		fmt.Println("Running in goroutine - new way")
	})

	// wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines finished execution.")
}
