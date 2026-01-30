package commuproblem

import "time"

func Sender(msg *string) {
	time.Sleep(500 * time.Millisecond)
	*msg = "Hello from Sender"
}

func Receiver(msg *string) {
	println(*msg)
}

func DelayedReceiver(msg *string) {
	time.Sleep(1 * time.Second)
	println(*msg)
}
