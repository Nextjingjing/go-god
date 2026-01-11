# บทที่ 1 Package

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO

## Package คืออะไร?
Package คือ วิธีการจัดการและการ **reuse code**  ใน go ซึ่งทำให้เราแยกโค้ดออกเป็นส่วนๆ ตามหน้าที่ที่เรากำหนด

## Package Manager
คือ ตัวจัดการ Package และเป็นระบบจัดการไลบรารีและ Dependencies ของโปรเจกต์
``` bash
$ go mod init <module-path>
$ go mod init github.com/<username>/... # example
```

## สร้าง package
สร้างโฟลเดอร์ `./hello/` และทำการสร้างไฟล์ `./hello/hello.go` 
``` go
package hello

func Greet() {
	println("Hello, package!")
}
```

**ต้องประกาศชื่อ `packe hello` package ให้ตรงกับชื่อโฟล์เดอร์ด้วย !** และอีกสิ่งสำคัญคือ `func Greet()` ต้อง G ใหญ่ด้วย ลองไปดูตัวอย่างการเรียกใช้ใน `./main.go` ดู ต่อมาเราสร้าง `./hello/greet.go`
```go
package hello

func Greet2() {
	println("hello, Greet2")
}
```
สังเกตว่า `packe hello` เหมือนกัน ไฟล์ใดๆที่อยู่ในโฟลเดอร์ `./hello/` จะใช้ Package นี้หมด

## Private และ Public
ที่ไฟล์ `./hello/hello.go` 
```go
package hello

func privateFunction() {
	println("Hello, private function!")
}

func PublicFunction() {
	println("Hello, public function!")
}
```

ดังตัวอย่างเลยการ Public ให้ Package อื่นไปใช้ได้ตัวแรกต้องพิมพ์ใหญ่เช่น `PublicFunction` ใช้ P ใหญ่ ในทางกลับกันอยากใช้ภายใน Package ต้องขึ้นต้นพิมพ์เล็กเช่น `privateFunction` ก็จะไม่สามารถใช้นอก Package ได้ ดูการเรียกใช้งานที่ `./main.go`

## Nested Packages (Package ซ้อน Package)
ที่ไฟล์ `./parent/child/chid.go` จะมีโค้ดคือ
```go
package child

func ChildFunc() {
	println("Hello Child")
}
```
ยังไงชื่อ Package ก็ดูแค่โฟลเดอร์ที่อยู่บนหัวมัน และตัวอย่างการใช้งานใน `./main.go`
```go
package main

import (
	"github.com/Nextjingjing/01-package/go-god/hello"
	"github.com/Nextjingjing/01-package/go-god/parent/child"
)

func main() {
	...

	child.ChildFunc()
}
```
และยังไงมันก็เป็น Package `child` คนละอันกับ `parent`