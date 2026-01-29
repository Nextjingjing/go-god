# Part 1 Goroutine

## Goroutine คืออะไร?
Goroutine คือ เป็นเหมือนการสร้าง Process/Thread ออกมาแยกจากการรันโปรแกรมหลัก (เหมือนมีหลายๆ โปรแกรมกำลังทำงานอยู่พร้อมกัน) เพื่อให้สามารถทำงานได้แบบพร้อมๆ กัน (Concurrently) ซึ่งบางครั้งโค้ดของเราอยากให้มีการทำงานอย่างพร้อมกัน 

เช่น Web server เมื่อมีหลายๆ Request เข้ามาหากเราเขียนแบบไม่ใช้ Goroutine ผลที่ได้คือเมื่อมีผู้ใช้เข้ามาหลายๆ คนพร้อมกัน แล้วทำให้การตอบสนองต้องรอคนแรกเสร็จ จากนั้นจึงจะทำคนที่สอง ...

## การใช้งาน Goroutine 
การใช้งาน Goroutine นั้นไม่ยากเลยเพียงแค่ใช้ `go` นำหน้า `func` ที่เราอยากให้ทำงานพร้อมๆ กัน อารมณ์จะเหมือนการสร้าง Thread ขึ้นมา แต่ขอบอกตรงนี้เลยว่า ***goroutine ไม่ใช่ thread*** แต่เป็นงานที่ `GO runtime` ต้องแมพงานนี้ไปหา `OS Thread` นั้นเอง
```go
package main

import "time"

func runner(msg string) {
	println(msg)
}

func main() {
	go runner("runner no. 1")
	go runner("runner no. 2")
	go runner("runner no. 3")

	time.Sleep(time.Second * 2) 
    // wait for all runners to finish or timeout
}

```
- function `runner(...)` จะถูกรันใน Thread พร้อมๆ กันทำให้เปรียบเสมือนการวิ่งพร้อมๆ กัน
- ***ผลลัพธ์ไม่แน่นอน*** เนื่องจากแต่ละ Goroutine จะทำงานพร้อมๆ กันและไม่อาจรู้ว่าใครจะทำเสร็จก่อน
```bash
runner no.2
runner no.1
runner no.3
```
หรือ
```bash
runner no.2
runner no.3
runner no.1
```
หรืออื่นๆ สามารถเกิดได้หมดขึ้นอยู่กับว่าใครจะได้ทำงานเสร็จก่อนกัน
- ผมจะมอง `go runner("runner no. 1")`, 
	`go runner("runner no. 2")` และ
	`go runner("runner no. 3")` เป็นโปรแกรมลูกๆ
- มอง `func main()` เป็นโปรแกรมแม่หรือโปรแกรมหลัก
- ข้อสังเกตต่อมาคือ `time.Sleep(time.Second * 2)` ทำไมโปรแกรมหลัก (โปรแกรมแม่) ต้องรอ 2 วินาที นั้นก็เพราะหากโปรแกรมแม่ทำงานเสร็จแล้ว มันจะจบโปรแกรมทันที ทำให้ลูกๆ ต้องจบโปรแกรมตามนั้นเอง (แม่ฆ่าลูกๆ ถ้าลูกเสร็จทีหลัง)

ให้คุณเปิดไปที่ไฟล์ `./main.go`
```go
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
```
- ลองรันดู
- คุณจะพบว่า `go slowerRunner(...)` จะไม่แสดงผลเลย เพราะเขาทำงาน 99 วินาที ซึ่งนานเกินเวลาของโปรแกรมแม่ที่รอแค่ 2 วินาที 