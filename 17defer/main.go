package main

import "fmt"

func main() {
	// fmt.Println("Go lang defer") // LIFO

	// defer fmt.Println("Hello")
	// fmt.Println("world")
	// the keyword 'defer' puts the statement at the end of function in reverse order - LIFO

	// output --
	// Go lang defer
	// world
	// Hello

	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Three")
	Counter()
}

func Counter() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// regular output flow
// One, Two Three
// 01234

// with defer output
// Three
// 43210Two
// One
