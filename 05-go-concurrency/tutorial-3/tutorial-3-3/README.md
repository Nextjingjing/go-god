# part 3.3 Sync package Condition Variable

## Condition Variable คืออะไร?
ในบางคร้ังเราต้องการจะรอจนกว่าจะมีเหตุการณ์บางอย่างที่ตรงกับเงื่อนไขที่เราต้องการ ดังนั้นแล้วทางแก้คืออะไรละ `While loop` หรอ
```go
for !condition {
    // Busy wait
}
```

เหมือนจะดีๆ แต่! `Busy waiting` การรอด้วย `While loop` ทำให้ CPU ไม่เกิดประโยชน์แทนที่จะเอาทรัพยากรไปทำอย่างอื่น 

`Condition Variable` จึงมาแก้จุดนี้

## การประกาศตัวแปร

```go
var mu sync.Mutex
cond := sync.NewCond(&mu)
```

## การสร้างเงื่อนไขรอ

```go
cond.L.Lock()
for !condition {
		cond.Wait()
	}
```

- การ `Wait()` เปรียบเสมือนการสั่งให้ Goroutine นั้นไปนอนรอ (ไม่ Busy waiting) และรอคนมาปลุก

## การปลุก Goroutine อื่นมาใช้งาน

```go
cond.Signal()
cond.L.Unlock()
```

- ปลุก Goroutine สักอันหนึ่งมาทำงาน
- `Signal()` จะปลุกแค่ 1 Goroutine เท่านั้นจากบรรดา Goroutine ที่ยัง `Wait()`


## ตัวอย่างเช่น 
`John` และ `Jane` ทั้งคู่ต่างอยากใช้ห้องน้ำ `Toilet` แต่ปัญหาคือมีได้แค่ 1 ห้องน้ำให้คนทั้งสองต้องรอเงื่อนไขว่ารอห้องน้ำ `available` ก่อนจึงจะสามารถเข้าได้

```go
package main

import (
	"sync"
	"time"
)

type People struct {
	name string
}

type Toilet struct {
	available bool
	cond      *sync.Cond
}

func (p *People) UseToilet(t *Toilet, wg *sync.WaitGroup) {
	defer wg.Done()

	t.cond.L.Lock()
	for !t.available {
		t.cond.Wait()
	}

	t.available = false
	time.Sleep(500 * time.Millisecond) // Simulate time taken to use the toilet
	println(p.name, "is using the toilet")

	t.available = true
	t.cond.Signal()
	t.cond.L.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	toilet := Toilet{
		available: true,
		cond:      sync.NewCond(&mu),
	}

	john := People{name: "John"}
	jane := People{name: "Jane"}

	wg.Add(2)
	go john.UseToilet(&toilet, &wg)
	go jane.UseToilet(&toilet, &wg)

	wg.Wait()
}

```
ผลลัพธ์
```bash
Jane is using the toilet
John is using the toilet
```
หรือ
```bash
John is using the toilet
Jane is using the toilet
```