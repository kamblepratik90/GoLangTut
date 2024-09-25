package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Go lang goroutine")

	// go greeter("Hello")
	// greeter("World")

	websites := []string{
		"https://www.google.com/",
		"https://www.fb.com",
		"https://www.x.com",
		"https://www.instagram.com",
	}
	for _, site := range websites {
		go getStatusCode(site)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(msg string) {

// 	for i := 0; i < 6; i++ {
// 		time.Sleep(5 * time.Millisecond)
// 		fmt.Println(msg)
// 	}

// }

func getStatusCode(endpoint string) int64 {
	defer wg.Done()
	_, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("\n oops, err for : %s", endpoint)
		return 0
	} else {
		fmt.Printf("\n 200 success for : %s", endpoint)
		return 200
	}
}
