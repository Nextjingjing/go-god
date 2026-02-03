package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		slowTask(ctx)
	}()

	go func() {
		defer wg.Done()
		fastTask(ctx)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("--- Ordering Cancellation ---")
	cancel()

	wg.Wait()
	fmt.Println("All goroutines exited safely.")
}

func slowTask(ctx context.Context) {
	fmt.Println("SlowTask: Started...")
	select {
	case <-time.After(99 * time.Second):
		fmt.Println("SlowTask: Finished successfully.")
	case <-ctx.Done():
		fmt.Println("SlowTask: Received cancel signal, stopping...")
	}
}

func fastTask(ctx context.Context) {
	fmt.Println("FastTask: Started...")
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("FastTask: Finished successfully.")
	case <-ctx.Done():
		fmt.Println("FastTask: Received cancel signal, stopping...")
	}
}
