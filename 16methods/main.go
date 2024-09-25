package main

import "fmt"

func main() {

	fmt.Println(" methods in GO lang")

	user1 := User{"Berlin", "a@b.com", 111, true}

	fmt.Println("myUser: ", user1) // {Berlin a@b.com 111 true}

	fmt.Printf("my user details %+v \n", user1) // {Name:Berlin Email:a@b.com Age:111 Status:true}

	fmt.Println("myUser Name: ", user1.Name) // Berlin

	// ----------------- methods
	status := user1.GetStatus()
	fmt.Println("myUser GetStatus: ", status) // actual status - true / false

	// IMP
	newEmail := user1.GetNewEmail()
	fmt.Println("myUser GetNewEmail: ", newEmail) // "zzzz@qqqq.com"

	// NOTE:
	fmt.Println("myUser Email: ", user1.Email) // a@b.com
	// Notice here, the user1 still returned a@b.com, this is cause the copy was been created, instead of
	// actual objet user1
	// Thats why to resolve this kind of things, we need to use POINTER - address ref

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() bool {
	return u.Status
}

func (u User) GetNewEmail() string {
	return "zzzz@qqqq.com"
}
