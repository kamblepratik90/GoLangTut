package main

import "fmt"

func main() {

	fmt.Println(" Go lang loops")

	days := []string{"Mon", "Tue", "Wed", "Thur", "Fri", "Sat", "Sun"}

	fmt.Println(" all days: ", days)

	// for
	for d := 0; d < len(days); d++ {
		fmt.Printf("d is %v and day is: %v \n", d, days[d])
	}
	fmt.Printf("----------------------------------------------\n")
	// for
	for i := range days {
		fmt.Println("day[i]: ", days[i])
	}

	// for
	fmt.Printf("----------------------------------------------\n")
	// for
	for index, day := range days {
		fmt.Printf(" index is : %v and day is : %v \n", index, day)
	}

	// for
	fmt.Printf("----------------------------------------------\n")
	// for
	for _, day := range days {
		fmt.Printf(" day is : %v \n", day)
	}

	// for
	fmt.Printf("----------------------------------------------\n")
	counter := 1
	for counter < 10 {
		fmt.Println(" counter is : ", counter)
		counter++
	}
}
