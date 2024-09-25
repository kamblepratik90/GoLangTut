package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Golang files operations")

	content := "This is the file content"

	fileName := "./myFile.txt"
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("content length : ", length)
	defer file.Close()

	readFile("./myFile.txt")

}

func readFile(filepathName string) {

	// databytes, err := ioutil.ReadFile(filepathName)
	databytes, err := os.ReadFile(filepathName)

	if err != nil {
		panic(err)
	}

	fmt.Println("the content of d the file is : \n ", databytes)

	fmt.Println("the content of d the file is : \n ", string(databytes))
}
