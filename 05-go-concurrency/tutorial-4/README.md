# Part 3 Deadlock

## Deadlock คืออะไร?
มันคือกลไกลป้องกันการรอแบบไม่รู้จบที่ อันเกิดมาจากการผิดพลาดในการเขียนโค้ด

### Deadlock เกิดตอนไหนบ้าง
1. รับค่าจาก channel แต่ไม่มีใครส่ง
```go
ch := make(chan int)
<-ch // รอรับ แต่ไม่มี goroutine ไหนส่ง
```
2. ส่งค่าเข้า channel แต่ไม่มีใครรับ (โดยเฉพาะ unbuffered channel)
```go
ch := make(chan int)
ch <- 10 // บล็อกตลอด เพราะไม่มีคนรับ
```
3. ใช้ mutex lock ซ้อน แล้วไม่ unlock
```go
mu.Lock()
mu.Lock() // lock ซ้ำ ตัวเองบล็อกตัวเอง
```
4. WaitGroup แล้วจำนวน Add กับ Done ไม่ตรง
```go
wg.Add(1)
go func() {
    // ลืม wg.Done()
}()
wg.Wait() // รอตลอดไป
```
5. channel เต็ม (buffered channel) แล้วส่งเพิ่ม
```go
ch := make(chan int, 1)
ch <- 1
ch <- 2 // เต็มแล้ว ไม่มีคนรับ = deadlock
```