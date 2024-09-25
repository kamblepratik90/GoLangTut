package main

import "fmt"

func main() {
	fmt.Println(" Go lang maps")

	myMap := make(map[string]string) // (map[KEY]value)
	fmt.Printf(" type id myMaps : %T \n", myMap)
	fmt.Println(" my map: ", myMap) // map[]

	myMap["JS"] = "JavaScript"
	myMap["JW"] = "Java"
	myMap["PY"] = "Python"
	myMap["RB"] = "Ruby"

	fmt.Println(" my map: ", myMap)           // map[JS:JavaScript JW:Java PY:Python RB:Ruby]
	fmt.Println(" JS is for : ", myMap["JS"]) // JavaScript

	// delete

	delete(myMap, "PY")
	fmt.Println(" my map: ", myMap) // map[JS:JavaScript JW:Java RB:Ruby]

	// loops thr map

	for key, value := range myMap {
		fmt.Printf(" the key : %v is for value : %v \n", key, value)
	}
}
