package main

import "time"

func runner(msg string) {
	println(msg)
}

func slowerRunner(msg string) {
	time.Sleep(time.Second * 99)
	println(msg)
}

func main() {
	println("Running race Start !!!")
	println("all runner will run concurrently")
	println("====================")
	go slowerRunner("slower runner")
	go runner("runner no. 1")
	go runner("runner no. 2")
	go runner("runner no. 3")
	go runner("runner no. 4")
	go runner("runner no. 5")
	go runner("runner no. 6")

	time.Sleep(time.Second * 2) // wait for all runners to finish or timeout
	println("timeout !!!")
	println("====================")
	println("slower runner did not finish in time!")
	// slowerRunner may not finish before main ends
}
