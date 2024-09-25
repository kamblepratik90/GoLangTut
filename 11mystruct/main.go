package main

import "fmt"

func main() {

	fmt.Println(" struct in GO lang")

	user1 := User{"Berlin", "a@b.com", 111, true}

	fmt.Println("myUser: ", user1) // {Berlin a@b.com 111 true}

	fmt.Printf("my user details %+v \n", user1) // {Name:Berlin Email:a@b.com Age:111 Status:true}

	fmt.Println("myUser Name: ", user1.Name) // Berlin
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
