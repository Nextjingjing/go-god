# Part 2.1 go channel

## ความเข้าใจ Thread
โดยทั่วไปแล้ว Thread นั้นจะแชร์ Variable ร่วมกัน แต่ใช้แยก Program counter (PC) กัน

![geeksforgeeks](https://media.geeksforgeeks.org/wp-content/uploads/20250829104433669224/multithreading-in-os.png) 

รูปจาก https://media.geeksforgeeks.org/wp-content/uploads/20250829104433669224/multithreading-in-os.png

```go
package main

import "time"

var n int = 42

func printAddress() {
	println(&n)
}

func main() {

	go printAddress()
	go printAddress()

	// wait for goroutines to finish
	time.Sleep(1 * time.Second)
}
```

โค้ดนี้ปริ้นท์ Address ของตัวแปร `n` ปรากฏว่าได้ผลลัพธ์ดังนี้
```bash
0x7ff68e915368
0x7ff68e915368
```
ซึ่งคือ Address เดียวกัน

## ปัญหาการสื่อสารระหว่าง Goroutine
บางครั้งเราอยากให้มีการสื่อสารระหว่าง `Goroutine` แต่มันก็ใช้ `Shared Variables` ในการสื่อสารกันได้ แต่ก็จะลำบากมากเพราะมันไม่มีการ `Synchronization` แบบกำหนดจังหวะการสื่อสารได้ 

เช่น ที่ไฟล์ `commu-problem/problem.go`
```go
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
```
- `Sender(msg *string)` ไว้ส่งข้อมูลผ่าน `Shared variables`
  - แต่จะเห็นว่าเรามีการ Delay ให้เขียนช้าลง 500 milliseconed
- `Receiver(msg *string)` ไว้อ่านข้อมูลผ่าน `Shared variables`
  - อ่านข้อมูลเลยทันที ไม่สนใจแม้ Sender จะส่งเสร็จยัง
- `DelayedReceiver(msg *string)` ไว้อ่านข้อมูลผ่าน `Shared variables`
  - รอ 1 วินาทีเพื่อให้มั่นใจว่า sender จะเขียนเสร็จ

การทดลองไปที่ไฟล์ `./main.go`

```go
package main

import (
	"time"

	commuproblem "example.com/commu-problem"
)

...

var msg string = "No message yet"

func main() {

	...

	time.Sleep(1 * time.Second)

	// simulate communication problem
	go commuproblem.Sender(&msg)
	go commuproblem.Receiver(&msg)
	go commuproblem.DelayedReceiver(&msg)

	// wait for goroutines to finish
	time.Sleep(3 * time.Second)
}

```
จะได้ผลลัพธ์คือ
```bash
No message yet
Hello from Sender
```
- ข้อความยังเขียนไม่เสร็จแต่คนอ่าน เลือกจะอ่านแล้ว

### นี้คือปัญหาเรื่องการสื่อสารระหว่าง Goroutine โดยไม่มีการ `Synchronization` 