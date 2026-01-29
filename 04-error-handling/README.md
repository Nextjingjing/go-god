# บทที่ 4 Error Handling

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO

## Error Handling คืออะไร?
Error handling คือการจัดการกับ Error แต่ขอบอกไว้ก่อนกนะครับว่า GO ไม่มี Try-Catch !!! แล้วจะมีวิธีจัดการกับ Error ยังไงเราไปชมกัน

## การ Throw Error ใน GO
```go
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}
```
จากตัวอย่างนี้จะเห็นว่าเราทำ `func Squrt(f float64)` และให้เขา return `(float64, error)` เจ้าตัว Error เป็น Type ของ Error นั้นเอง
- `errors.New("math: square root of negative number")` คือการสร้าง Error ขึ้นมาและสามารถใส่ Message ข้อมูลเกี่ยวกับ Error ลงไปได้นั้นเอง
- หากไม่มี Error เราจะ `return math.Sqrt(f), nil` ส่ง `nil` กลับไป เพราะไม่มี Error

## การรับมือกับ Error และการ Logging
```go
func main() {
	v, err := Sqrt(12)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v)
    ...
}
```
จากตัวอย่างนี้จะเห็นว่าเรารับมือโดยการเช็คว่า `err != nil` คือถ้ามี Error เราจะ `log.Fatal(err)` หรือการสั่งหยุดโปรแกรมนั้นๆ
```go
// Example of handling error without terminating the program
v, err = Sqrt(-12)
	if err != nil {
		log.Println(err)
	}
``` 
`log.Println(err)` แบบนี้ถ้า Error จะไม่หยุดโปรแกรม

```go
// Example of handling error by terminating the program
v, err = Sqrt(-12)
if err != nil {
	log.Fatal(err)
}
```
`log.Fatal(err)` แบบนี้ถ้า Error จะหยุดโปรแกรมทันที

```bash
2026/01/29 16:32:47 math: square root of negative number
exit status 1
```