package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26&ab_channel=HiteshChoudhary"

func main() {

	fmt.Println("Go lang url parsing")

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("the url props as below: ")
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	qParam := result.Query()
	fmt.Println(qParam)

	// v, list, index, ab_channel
	fmt.Println(qParam["v"])
	fmt.Println(qParam["list"])
	fmt.Println(qParam["index"])
	fmt.Println(qParam["ab_channel"])
	fmt.Println()

	for _, value := range qParam {
		fmt.Println("value : ", value)
	}

}
