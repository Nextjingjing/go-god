# บทที่ 2 Struct

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO

## Struct คืออะไร?
Struct ใน GO คือเป็นชนิดของข้อมูลประเภทหนึ่งที่สามารถ `Encapsulate` Field/Attribute และสามารถผูก Method ได้คล้ายๆ `Class` ของภาษาที่มี OOP นั้นเอง

## การสร้าง Struct 
ให้คุณเปิดไปที่ไฟล์ `./user/user.go` คุณจะเห็นมีโค้ดส่วนหนึ่งคือ
```go
package user

type User struct {
	id    int
	Name  string
	Email string
}
```
นี้คือวิธีการประกาศโครงสร้างของ Struct ของเราว่าต้องมี Field อะไรบ้าง?

### การสร้าง Instance ของ Struct
```go
func main() {
	u := user.User{
		Name:  "Nextjingjing",
		Email: "next@example.com",
	}
    ...
}
```
และสามารถเข้าถึง Attribute ได้โดยการ `instance.attributeName`
```go
func main() {
    ...
	println("User Name:", u.Name)
	println("User Email:", u.Email)
    ...
}
```

## การสร้าง Method ไปผูกกับ Struct
```go
func (u *User) GetID() int {
	return u.id
}

func (u *User) SetID(id int) {
	u.id = id
}

func NewUser(id int, name string, email string) User {
	return User{
		id:    id,
		Name:  name,
		Email: email,
	}
}
```
- Method `GetID` เป็น Getter ของค่า id ที่เป็น Private
- Method `SetID` เป็น Setter ของค่า id ที่เป็น Private
- Method `NewUser` เป็นคล้าย Constructor ในภาษา OOP ใช้สร้าง Instance ของ `User`

### การเรียกใช้ Method
```go
u2 := user.NewUser(2, "Gopher", "gopher@example.com")
id2 := u2.GetID()
println("User ID:", id2)
println("User Name:", u2.Name)
println("User Email:", u2.Email)
```
หรือ
```go
u3 := user.User{}
u3.SetID(3)
u3.Name = "John Doe"
u3.Email = "john@example.com"
id3 := u3.GetID()
println("User ID:", id3)
println("User Name:", u3.Name)
println("User Email:", u3.Email)
```