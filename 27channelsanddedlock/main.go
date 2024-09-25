package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("go lang - channel and dead lock")

	// init - declare channel
	// myChan := make(chan int)
	myChan := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// add value to channel with special operator
	// myChan <- 5

	// read from channel
	// fmt.Println("myChan: ", <-myChan) // fatal error: all goroutines are asleep - deadlock!

	wg.Add(2)
	// write in channel
	// go func(myChan chan int, wg *sync.WaitGroup) {
	go func(myChan chan<- int, wg *sync.WaitGroup) { // notice 'chan<-'
		fmt.Println("write chan")
		myChan <- 5
		myChan <- 4
		close(myChan)
		wg.Done()
	}(myChan, wg)

	// read from channel
	// go func(myChan chan int, wg *sync.WaitGroup) {
	go func(myChan <-chan int, wg *sync.WaitGroup) { // notice '<-chan'
		fmt.Println("read chan")
		fmt.Println("myChan: ", <-myChan)

		value, isChanelOpen := <-myChan

		fmt.Println("isChanelOpen: ", isChanelOpen)

		if isChanelOpen {
			fmt.Println("myChan value: ", value)
		}

		wg.Done()
	}(myChan, wg)

	// write chan
	// read chan
	// myChan:  5

	wg.Wait()
}
