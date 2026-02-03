# บทที่ 6 Context

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO
- บทที่ 5 Go concurrency

## Context คืออะไร? 
context คือ แพ็คเกจที่ช่วยให้สามารถจัดการการทำงานของ `Goroutine` ช่วยให้สามารถจัดการ `Deadline/Timeout`, `Cancellation` และสุดท้ายยังช่วย `Values Passing` 

## `context.Background()` และ `context.TODO()`
ใช้สร้าง Parent Context ซึ่งเป็น Context ว่างๆ `(Empty Context)` ที่ยังไม่กำหนด `Deadline/Timeout`, `Cancellation` และ `Values Passing` เลย โดยมีจุดต่างกันดังนี้
### 1. context.Background()
- Empty Context เอาไว้สร้าง Life cycle ของ Goroutine
- ใช้บ่อย 99%
- มั่นใจ แน่ใจว่าจะ Empty Context ไปสร้าง Context `Deadline/Timeout` หรือ `Cancellation` หรือ `Values Passing` แน่นอน

### 2. context.TODO()
- Empty Context เอาไว้สร้าง Life cycle ของ Goroutine
- ไม่มั่นใจว่าจะใช้ไปสร้าง `Deadline/Timeout` หรือ `Cancellation` หรือ `Values Passing` หรือไม่ ?
- ไม่ควรใช้จริง คำว่า `TODO` หมายถึงต้องมาแก้ 

### แต่ทั้งสองได้ Empty Context เหมือนกัน แต่แค่จุดประสงค์ต่างๆกัน

## การประกาศ Parent Context
```go
parentCtx := context.Background()
```
หรือ
```go
parentCtx := context.TODO()
```

## การสร้าง Child Context
เพื่อกำหนดว่าจะใช้ Feature `Deadline/Timeout` หรือ `Cancellation` หรือ `Values Passing` ได้

### Cancellation 
ใช้เพื่อยกเลิก Goroutine ที่ถือครอง `Child Context`
```go
childCtx, cancel := context.WithCancel(parentCtx)
```
หรือ 
```go
ctx, cancel := context.WithCancel(context.Background())
```
- ซึ่งแบบนี้จะได้ `cancel()` ซึ่งเป็น function ที่ช่วยให้ในการสั่ง Cancel เจ้าตัว Goroutine ที่ถือ `ctx` ได้

### Deadline/Timeout
ใช้เพื่อกำหนดเวลาตาย Goroutine ที่ถือครอง `Child Context`
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
```
- ซึ่งแบบนี้จะได้ `cancel()` ซึ่งเป็น function ที่ช่วยให้ในการสั่ง Cancel เจ้าตัว Goroutine ที่ถือ `ctx` ได้

### Values Passing
ใช้ส่ง Value ให้ Goroutine ที่ถือครอง `Child Context`
```go
ctx := context.WithValue(context.Background(), "userID", 1234)
```
และ Goroutine ใดๆ สามารถเอาค่า `1234` ได้โดย
```go
// Goroutine ต้องถือครอง ctx ก่อนนะ
userID := ctx.Value("userID").(int)
```
- Goroutine ต้องถือครอง `ctx` ก่อนนะ ถึงจะอ่านได้

## การถือครอง Child Context
```go
func worker(ctx context.Context, ch chan string) {
	// some work
}
```
- `ctx context.Context` ต้องเป็น Parameter แรกเท่านั้น

### การใช้ Select case ใน func ที่ถือครอง Child Context
```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            // เมื่อมีการเรียก cancel() หรือ Timeout
            fmt.Println("Worker: หยุดทำงานเพราะ", ctx.Err())
            return
        default:
            // ทำงานหลักต่อไป
            fmt.Println("Worker: กำลังประมวลผล...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
```

## ตัวอย่างเช่น

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		slowTask(ctx)
	}()

	go func() {
		defer wg.Done()
		fastTask(ctx)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("--- Ordering Cancellation ---")
	cancel()

	wg.Wait()
	fmt.Println("All goroutines exited safely.")
}

func slowTask(ctx context.Context) {
	fmt.Println("SlowTask: Started...")
	select {
	case <-time.After(99 * time.Second):
		fmt.Println("SlowTask: Finished successfully.")
	case <-ctx.Done():
		fmt.Println("SlowTask: Received cancel signal, stopping...")
	}
}

func fastTask(ctx context.Context) {
	fmt.Println("FastTask: Started...")
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("FastTask: Finished successfully.")
	case <-ctx.Done():
		fmt.Println("FastTask: Received cancel signal, stopping...")
	}
}

```
ผลลัพธ์
```bash
SlowTask: Started...
FastTask: Started...
FastTask: Finished successfully.
--- Ordering Cancellation ---
SlowTask: Received cancel signal, stopping...
All goroutines exited safely.
```