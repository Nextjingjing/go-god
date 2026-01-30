package main

import (
	"time"

	commuproblem "example.com/commu-problem"
)

var n int = 42

func printAddress() {
	println(&n)
}

var msg string = "No message yet"

func main() {

	go printAddress()
	go printAddress()

	time.Sleep(1 * time.Second)

	// simulate communication problem
	go commuproblem.Sender(&msg)
	go commuproblem.Receiver(&msg)
	go commuproblem.DelayedReceiver(&msg)

	// wait for goroutines to finish
	time.Sleep(3 * time.Second)
}
