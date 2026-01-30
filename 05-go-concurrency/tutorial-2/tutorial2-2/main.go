package main

import (
	"fmt"
	"time"
)

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

	fmt.Printf("capacity of channel: %d\n", cap(ch))

	// wait for goroutines to finish
	time.Sleep(2 * time.Second)

	println("=================")
	println("Buffer Channel")
	bufferedCh := make(chan string, 4)
	fmt.Printf("capacity of bufferedCh: %d\n", cap(bufferedCh))
	fmt.Printf("size of bufferedCh: %d\n", len(bufferedCh))
	bufferedCh <- "msg 1"
	bufferedCh <- "msg 2"
	bufferedCh <- "msg 3"
	bufferedCh <- "msg 4"
	go reader(bufferedCh)
	go reader(bufferedCh)
	time.Sleep(1 * time.Second)
	fmt.Printf("size of bufferedCh: %d\n", len(bufferedCh))
}
