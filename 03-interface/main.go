package main

import (
	"fmt"

	"github.com/Nextjingjing/go-god/03-interface/shape"
	"github.com/Nextjingjing/go-god/03-interface/user"
)

func main() {
	fmt.Println("Shape Interface Example")
	var s shape.Shape

	s = shape.Circle{Radius: 5}
	fmt.Println("C Area:", s.Area())
	fmt.Println("C Perimeter:", s.Perimeter())

	s = &shape.Rectangle{Length: 4, Width: 3}
	fmt.Println("R Area:", s.Area())
	fmt.Println("R Perimeter:", s.Perimeter())

	// s = &shape.Triangle{Base: 4, Height: 3}
	/*
		cannot use &shape.Triangle{…} (value of type *shape.Triangle) as shape.Shape value in assignment:
		*shape.Triangle does not implement shape.Shape (missing method Perimeter)
	*/

	// Even though trapezoid has Additional method, it still implements the Shape interface
	s = &shape.Trapezoid{BaseTop: 3, BaseBottom: 5, Height: 4, SideLeft: 2, SideRight: 2}
	fmt.Println("T Area:", s.Area())
	fmt.Println("T Perimeter:", s.Perimeter())

	fmt.Println("======================")
	fmt.Println("User Repository Interface Example")
	// Using MockUserRepository
	mockRepo := &user.MockUserRepository{
		Users: map[int]user.User{
			1: {ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
			2: {ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
		},
	}

	userService := &user.UserService{Repo: mockRepo}
	u, err := userService.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User from Mock Repo: %+v\n", u)
	}

	// Using PostgresUserRepository
	pgRepo := &user.PostgresUserRepository{}
	userService.Repo = pgRepo
	u, err = userService.GetUser(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User from database Repo: %+v\n", u)
	}
}
