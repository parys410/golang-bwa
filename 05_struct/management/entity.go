package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// Method is a Struct Function
func (u User) Display() (result string) {
	result = fmt.Sprintf("Name : %s, Email : %s", u.FirstName + " " + u.LastName, u.Email)
	return
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func (g Group) Display() {
	fmt.Printf("\n%s's Group (%d Member)\n", g.Name, len(g.Users))
	fmt.Printf("Member : ")
	for _,member := range g.Users {
		fmt.Printf("%s ", member.FirstName + member.LastName)
	}
}