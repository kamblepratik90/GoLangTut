package main

func main() {

	println(" this is pointers from golang")

	var mypointer *int

	println("mypointer ", mypointer)

	myNumber := 32

	var newPointer = &myNumber

	println("newPointer location: ", newPointer) // address loc
	println("newPointer value: ", *newPointer)   // value

	*newPointer = *newPointer + 2

	println("new newPointer location: ", newPointer)
	println("new newPointer value: ", *newPointer)

}
