package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	welcome := "Welcome to my app"
	fmt.Println(welcome)
	fmt.Println("Enter your rating from 1 - 5")

	reader := bufio.NewReader(os.Stdin)

	// comma, ok || err ok

	// accept user input until user hit \n i.e enter
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of rating, %T", input)

	// trim input and parse - it accepts \n with input, we need to trim that
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("\nError is ,", err)
	} else {
		// we need to fo operations on numRating
		newNumRating := numRating + 1*1
		fmt.Println("Final rating is : ", newNumRating)
	}

}
