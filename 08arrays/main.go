package main

import "fmt"

func main() {

	fmt.Println("This is arrays.. lesser used in GO lang")

	var fruitArray [4]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Banana"
	// skip [2]
	fruitArray[3] = "Cherry"

	fmt.Println("fruitArray is ", fruitArray) // [Apple Banana  Cherry]
	fmt.Println("fruitArray length ", len(fruitArray))

	var vegList = [5]string{"potato", "beans", "tomato"}
	fmt.Println("vegList is ", vegList) // [potato beans tomato  ]
	fmt.Println("vegList length ", len(vegList))

}
