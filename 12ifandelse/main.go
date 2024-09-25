package main

import "fmt"

func main() {
	fmt.Println("Go lang if and else... control statements")

	count := 23

	var message string

	if count < 20 {
		message = "lesser count"
	} else if count > 20 {
		message = "higher count"
	} else {
		message = "count is exactly 20 "
	}

	fmt.Println(" message is : ", message)
}
