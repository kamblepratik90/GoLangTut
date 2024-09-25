package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://example.com/"

func main() {
	fmt.Println(" Go lang simple web req")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type : %T \n", response)

	defer response.Body.Close() // remember this is imp

	// dataBytes, err := ioutil.ReadAll(response.Body) // deprecated
	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println("the content of web API response : \n", content)

	// OR

	var responseString strings.Builder
	byteCount, _ := responseString.Write(dataBytes)
	fmt.Println("the content len : ", byteCount)
	fmt.Println("the content of web API response : ", responseString.String())
}
