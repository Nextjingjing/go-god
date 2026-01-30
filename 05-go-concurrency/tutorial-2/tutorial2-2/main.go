package main

import "time"

func writer(ch chan string, msg string) {
	time.Sleep(500 * time.Millisecond)
	println("writer send msg to another goroutine")
	ch <- msg
}

func reader(ch chan string) {
	// synchronization
	msg := <-ch // waits for a message from the channel
	println("reader received message:", msg)
	println("==================")
}

func main() {
	ch := make(chan string)
	go writer(ch, "hello channel 1")
	go reader(ch)

	go reader(ch)
	go writer(ch, "hello channel 2")
	// wait for goroutines to finish
	time.Sleep(2 * time.Second)
}
