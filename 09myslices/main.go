package main

import (
	"fmt"
	"sort"
)

func main() {
	println(" Welcome GOlang slices")

	// var fruitList = []string{} // init
	var fruitList = []string{"Apple", "Oranges", "Banana"} // init
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	// add - .append
	fruitList = append(fruitList, "Mango")

	// slice the slices
	// fruitList = fruitList[0:]

	fmt.Println("fruitList is ", fruitList)      // [Apple Oranges Banana Mango]
	fmt.Println("fruitList is ", fruitList[0:])  // [Apple Oranges Banana Mango]
	fmt.Println("fruitList is ", fruitList[1:4]) // [Oranges Banana, Mango]       // number non inclusive list
	fmt.Println("fruitList is ", fruitList[0:0]) // []
	fmt.Println("fruitList is ", fruitList[1:1]) // []
	fmt.Println("fruitList is ", fruitList[:0])  // []
	fmt.Println("fruitList is ", fruitList[:3])  // [Apple Oranges Banana]
	// fmt.Println("fruitList is ", fruitList[:10]) // - error
	// - panic: runtime error: slice bounds out of range [:10] with capacity 6

	// # syntax 2

	highScores := make([]int, 4)                 // slice having fixed size as 4
	fmt.Println("high scores are: ", highScores) // [0 0 0 0]

	highScores[0] = 123
	highScores[1] = 456
	highScores[2] = 789
	highScores[3] = 001
	fmt.Println("high scores are: ", highScores) // [123 456 789 1]

	// highScores[4] = 111
	// fmt.Println("high scores are: ", highScores) // error
	// - panic: runtime error: index out of range [4] with length 4

	highScores = append(highScores, 101, 102, 103) // but surprisingly this will be added
	fmt.Println("high scores are: ", highScores)   // [123 456 789 1 101 102 103]

	// check is sorted
	fmt.Println("high scored are sorted ?  -> ", sort.IntsAreSorted(highScores)) // true/false

	// sort
	sort.Ints(highScores)
	fmt.Println("sorted highScores: ", highScores) // [1 101 102 103 123 456 789]

	// check is sorted
	fmt.Println("high scored are sorted ?  -> ", sort.IntsAreSorted(highScores)) // true/false

	// remove value from slices based on index
	var courses = []string{"A", "B", "C", "D", "E"}
	fmt.Println("available courses are : ", courses) // [A B C D E]

	// now, need to remove course at index 2
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("available courses are : ", courses) // [A B D E]

}
