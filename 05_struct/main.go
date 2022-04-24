package main

import (
	"bwa-go/05_struct/management"
	"fmt"
)

func main() {
	user := management.User{
		ID: 1,
		FirstName: "Ary",
		LastName: "Setiyawan",
		Email: "parys410@gmail.com",
		IsActive: true,
	}

	user2 := management.User{}
	user2.ID = 2
	user2.FirstName = "Dwija"
	user2.LastName = "Shenda"
	user2.Email = "dwija@gmail.com"
	user2.IsActive = true

	user3 := management.User{ ID:3, FirstName: "Bangkit", LastName: "Aryawan", Email: "bangkit@gmail.com", IsActive: true}

	fmt.Println(user, user2, user3)
	fmt.Println(user.FirstName, user.LastName)
	fmt.Println(user3.Display())

	group1 := management.Group{
		Name: "Admin",
		Admin: user,
		Users: []management.User{ user, user2, user3 },
		IsAvailable: true,
	}
	group1.Display()
}

