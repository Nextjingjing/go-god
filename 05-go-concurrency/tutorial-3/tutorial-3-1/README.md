# part 3.1 Sync package WaitGroup

## WaitGroup คืออะไร?
โดยปกติแล้ว `Goroutine แม่` เมื่อทำงานเสร็จจะหยุดการทำงานโดยไม่สน `Goroutine ลูกๆ` ทำให้เราต้องใช้ `time.Sleep` เพื่อหน่วงเวลา แต่ก็มีปัญหาคือ
- ไม่รู้ว่าจะหน่วงกี่หน่วยเวลา จึงจะรับประกันว่าลูกๆ ทุกคนได้ทำครบแล้ว

```go
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
```
- `defer ...` คือ คำสั่งที่จะรับประกันว่าจะรันเมื่อทุกๆ คำสั่งๆ ใน function ที่ห่อหุ้มมันทำงานครบแล้ว (มันจะทำงานก่อน return คิดงี้)
- `var wg sync.WaitGroup` คือ การประกาศตัวแปร WaitGroup
- `wg.Add(1)` คือ การเพิ่มจำนวน Goroutine ที่ระบบต้องรอ เลข 1 หมายถึงเพิ่มมาหนึ่ง
- `defer wg.Done()` คือ บอก WaitGroup ว่ามี Goroutine เสร็จแล้วหนึ่งอัน
- `wg.Wait()` บอก WaitGroup ให้รอจนกว่าจะไม่มีงานเหลือแล้ว