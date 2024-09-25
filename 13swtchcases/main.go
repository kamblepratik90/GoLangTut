package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Go lang switch cases")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("dice number is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println(" it is 1")
	case 2:
		fmt.Println(" it is 2")
	case 3:
		fmt.Println(" it is 3")
	case 4:
		fmt.Println(" it is 4")
		fallthrough // this is special case - when 4 then also execute whatever is in next case
	case 5:
		fmt.Println(" it is 5")
	case 6:
		fmt.Println(" it is 6")
	default:
		fmt.Println(" anything else than 1-6, it is:  ", diceNumber)
	}
}
