package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello for time operations")

	presentTime := time.Now()
	fmt.Println("Present Time: ", presentTime)

	fmt.Println("Format date : ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.September, 1, 2, 20, 0, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
