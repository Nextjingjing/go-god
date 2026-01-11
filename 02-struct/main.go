package main

import "github.com/Nextjingjing/go-god/02-struct/user"

func main() {
	u := user.User{
		Name:  "Nextjingjing",
		Email: "next@example.com",
	}
	u.SetID(1)
	id := u.GetID()
	println("User ID:", id)
	println("User Name:", u.Name)
	println("User Email:", u.Email)

	u2 := user.NewUser(2, "Gopher", "gopher@example.com")
	id2 := u2.GetID()
	println("User ID:", id2)
	println("User Name:", u2.Name)
	println("User Email:", u2.Email)

	u3 := user.User{}
	u3.SetID(3)
	u3.Name = "John Doe"
	u3.Email = "john@example.com"
	id3 := u3.GetID()
	println("User ID:", id3)
	println("User Name:", u3.Name)
	println("User Email:", u3.Email)
}
