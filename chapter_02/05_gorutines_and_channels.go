package main

import (
	"fmt"
	"sync"
)

// Simulate processing data
func processData(data int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	// Simulate data processing with a simple operation
	result := data * 2
	results <- result
}
func main() {
	var wg sync.WaitGroup
	dataSets := []int{1, 2, 3, 4, 5}
	results := make(chan int, len(dataSets))
	for _, data := range dataSets {
		wg.Add(1)
		go processData(data, &wg, results)
	}
	// Close the results channel once all goroutines have finished
	go func() {
		wg.Wait()
		close(results)
	}()
	// Collect results
	for result := range results {
		fmt.Println(result)
	}
}