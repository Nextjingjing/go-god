# Part 2.2 go channel

จากปัญหาในครั้งก่อนนั้นคือการไม่มีการประสานงาน `Synchronization` ระหว่าง `Goroutine` ดังนั้นใน Part 2.2 นี้จะนำเสนอ `Channel` ตัวที่ช่วยให้ Goroutine คุยกันได้ง่าย

## Channel คืออะไร?
Channel คือ "ท่อส่งของ" ที่ทำให้ Goroutine ติดต่อสื่อสารกันเพื่อแบบปลอดภัยไม่ต้องมาประสานการทำงานเอง โดยมี 2 ประเภทของ Channel คือ
1. Unbuffer Channel
2. Buffer Channel

### 1. การสร้าง Unbuffer Channel
`Unbuffer Channel` คือ Channel ที่มีความจุของข้อมูลได้แค่ 1 อัน
```go
ch := make(chan string)
```
สร้าง `Channel` ที่ใช้ `string` สื่อสารกัน โดย Channel เปรียบเสมือนการท่อส่งข้อมูลที่ `Goroutine` ใดๆ สามารถนำข้อมูลมาใส่ และให้ Goroutine ตัวอื่นๆ มาอ่านได้ 

## การใส่ข้อมูลลงไปใน Channel
```go
ch <- "some data"
```

## การอ่านข้อมูลใน Channel
```go
data <- ch
```
- การอ่านข้อมูลนี้จะ

```go
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
```
- `writer()` ใช้ส่งข้อมูลไปใน `Channel` โดยทีมีการหน่วงเวลาก่อนเขียนด้วย
- `reader()` ใช้อ่านข้อมูลใน `Channel` 
เนื่องจากมีการหน่วงเวลาก่อนที่ `reader()` จะอ่านได้เราจะมาดูผลลัพธ์ว่าจะพังไหมหาก `reader()` ทำงานก่อน

ที่ไฟล์ `main.go`

```go
package main

import "time"

...

func main() {
	ch := make(chan string)
	go writer(ch, "hello channel 1")
	go reader(ch)

	go reader(ch)
	go writer(ch, "hello channel 2")
	// wait for goroutines to finish
	time.Sleep(2 * time.Second)
}

```
ผลลัพธ์
```bash
writer send msg to another goroutine
writer send msg to another goroutine
reader received message: hello channel 1
==================
reader received message: hello channel 2
==================
```
- จะสังเกตได้ว่าไม่ว่าจะลำดับ `writer()`, `reader()` ใดๆ จะไม่เกิดปัญหาอ่านก่อนเขียน เลยสักนิด เพราะ `Channel` ได้ประสานจังหวะให้อ่านเขียนได้ถูกต้องแล้ว
- กล่าวคือ `msg := <-ch // waits for a message from the channel` จะรอให้มี Goroutine ใดๆ เขียนลง `Channel`  นั้นๆ

### 2. การสร้าง Buffer Channel
`Unbuffer Channel` คือ Channel ที่มีความจุของข้อมูลได้ตามที่จองไว้ 

เช่น จอง 4 ช่อง

```go
package main

...

func main() {
	...

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
```
ผลลัพธ์
```bash
=================
Buffer Channel
capacity of bufferedCh: 4
size of bufferedCh: 0
reader received message: msg 1
==================
reader received message: msg 2
==================
size of bufferedCh: 2
```

## ข้อแตกต่าง Buffer vs Unbuffer Channel
| หัวข้อเปรียบเทียบ              | Unbuffered Channel            | Buffered Channel                |
| :------------------------- | :---------------------------- | :------------------------------ |
| **การประกาศ**              | `make(chan T)`                | `make(chan T, n)`               |
| **ความจุ (Capacity)**       | **0**                         | **n** (ตามที่ระบุ)                 |
| **ประเภทการส่ง**            | **Synchronous** (ส่ง-รับพร้อมกัน) | **Asynchronous** (ส่งทิ้งไว้ได้)     |
| **พฤติกรรมผู้ส่ง (Sender)**    | จะ Block จนกว่าจะมีคนมา `<-ch`  | จะ Block เมื่อ **Buffer เต็ม**     |
| **พฤติกรรมผู้รับ (Receiver)**  | จะ Block จนกว่าจะมีคนมา `ch <-` | จะ Block เมื่อ **Buffer ว่าง**     |

### การ Block คือการที่ Goroutine จะไม่ถูกทำงานจนกว่าจะเข้าเงื่อนไขบางอย่าง
เช่น `reader(ch)` ผู้อ่านจะถูก Block ไม่ให้ทำงานต่อหาก `writter(ch, "hello")` ยังไม่เขียนเสร็จสังเกตว่าโค้ดผมจะหน่วงเวลาเขียนไว้ด้วย