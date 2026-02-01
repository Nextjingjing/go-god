# part 3.2 Sync package Mutex

## Mutex คืออะไร?
Mutex คือกลไกลป้องกันการทำงานในส่วน Critical section พร้อมๆ กันทำให้ทำงานผิดพลาดได้

Critical section คือโค้ดส่วนที่อันตรายหากถูกทำงานพร้อมๆ กัน

```go
mutex.Lock()
// Critical section
mutex.Unlock()
```

ตัวอย่างเช่น 

```go
package main

import "sync"

type counter struct {
	mutex sync.Mutex
	val   int
}

func (c *counter) incrementBy1000000(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mutex.Lock()
	for i := 0; i < 1000000; i++ {
		c.val++
	}
	c.mutex.Unlock()
}

func (c *counter) incrementBy1000000NoMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		c.val++
	}
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	c_mutex := counter{
		val:   0,
		mutex: mutex,
	}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go c_mutex.incrementBy1000000(&wg)
	}
	wg.Wait()
	println("c_mutex.val:", c_mutex.val)

	c_no_mutex := counter{
		val:   0,
		mutex: mutex,
	}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go c_no_mutex.incrementBy1000000NoMutex(&wg)
	}
	wg.Wait()
	println("c_no_mutex.val:", c_no_mutex.val)
}
```
ผลลัพธ์
```bash
c_mutex.val: 200000
c_no_mutex.val: 200000
```
ผลลัพธ์
```bash
c_mutex.val: 200000
c_no_mutex.val: 113245
```
ผลลัพธ์
```bash
c_mutex.val: 2000000
c_no_mutex.val: 1074882
```

## การประกาศตัวแปร Mutex
```go
var mutex sync.Mutex
```

## การใช้งาน
```go
mutex.Lock()
// Critical section
mutex.Unlock()
```
