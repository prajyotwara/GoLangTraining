package main

import (
	"fmt"
	"sync"
)

func ConcurrencyWithWaitGroupExample() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(count int) {
			fmt.Println("Running concurrently : ", count)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Completed")
}

func ConcurrencyWithWaitGroupExampleUsingDefer() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()
			fmt.Println("Running concurrently : ", count)

		}(i)
	}
	wg.Wait()
	fmt.Println("Completed")
}
