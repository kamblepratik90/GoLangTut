package main

import "fmt"

func main() {
	fmt.Println(" Go lang functions...")

	c := adder(2, 3)
	fmt.Println("adder value : ", c)

	d := newAdder(1, 2, 3, 4, 5)
	fmt.Println("new adder value : ", d)

	value, msg := newAdderTwo(10, 2, 3, 4, 5)
	fmt.Println("new newAdderTwo value : ", value)
	fmt.Println("new newAdderTwo message : ", msg)

}

func adder(a int, b int) int {
	return a + b
}

func newAdder(mySlice ...int) int {
	total := 0

	for _, value := range mySlice {
		total += value
	}

	return total
}

func newAdderTwo(mySlice ...int) (int, string) {
	total := 0

	for _, value := range mySlice {
		total += value
	}

	return total, "hello msg"
}
