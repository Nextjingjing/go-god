package main

import "sync"

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	println(message)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printMessage("Hello from goroutine", &wg)
	}
	wg.Wait()
	// instead of time.Sleep, we use WaitGroup to wait for all goroutines to finish
}
