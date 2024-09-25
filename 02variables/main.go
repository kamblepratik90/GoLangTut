package main

import "fmt"

// const
const IamConst string = "aaaaaaaaa" // public

func main() {
	fmt.Println("hello variables")
	var username string = "userBingo"
	fmt.Println("\nhello variables", username)
	fmt.Printf("variable is of type: %T \n", username)

	var isActive bool = true
	fmt.Println("\nhello variables", isActive)
	fmt.Printf("variable is of type: %T \n", isActive)

	var smallVal uint8 = 255
	fmt.Println("\nhello variables", smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float32 = 25678.5465656898089
	fmt.Println("\nhello variables", smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// default value

	var myName string
	fmt.Println("\nhello variables", myName)
	fmt.Printf("variable is of type: %T \n", myName)

	// implicit type
	var yourName = "userBingo"
	fmt.Println("\nhello variables", yourName)
	fmt.Printf("variable is of type: %T \n", yourName)

	// no var type -> := walrus operator (can be declared inside method only)
	newName := "GOGO"
	fmt.Println("\nhello variables", newName)
	fmt.Printf("variable is of type: %T \n", newName)

	fmt.Println("\nhello variables", IamConst)
	fmt.Printf("variable is of type: %T \n", IamConst)

}
