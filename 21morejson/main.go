package main

import (
	"encoding/json"
	"fmt"
)

// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string
// 	Password string
// 	Tags     []string
// }

type course struct { // aliases
	Name     string `json:"courseName"` // mind the spaces
	Price    int
	Platform string   `json:"websites"`       // mind the spaces
	Password string   `json:"-"`              // used for hiding the field from JSON response
	Tags     []string `json:"tags,omitempty"` // remove empty (ni/null) records
}

func main() {
	fmt.Println("Go lang JSON")

	// EncodeJson()

	DecodeJson()
}

func EncodeJson() {
	coursesList := []course{
		{"ReactJs", 299, "YouTube", "123qwe", []string{"JS", "web"}},
		{"MERN", 199, "Tube", "abcqwe", []string{"JS", "full"}},
		{"Angular", 99, "You cube Tube", "123qwe", nil},
	}

	// package this as JSON data

	// response, err := json.Marshal(coursesList)
	// [{"Name":"ReactJs","Price":299,"Platform":"YouTube","Password":"123qwe","Tags":["JS","web"]},{"Name":"MERN","Price":199,"Platform":"Tube","Password":"abcqwe","Tags":["JS","full"]},{"Name":"Angular","Price":99,"Platform":"You cube Tube","Password":"123qwe","Tags":null}]

	response, err := json.MarshalIndent(coursesList, "", "\t")
	// 	[
	//         {
	//                 "Name": "ReactJs",
	//                 "Price": 299,
	//                 "Platform": "YouTube",
	//                 "Password": "123qwe",
	//                 "Tags": [
	//                         "JS",
	//                         "web"
	//                 ]
	//         },
	//         {
	//                 "Name": "MERN",
	//                 "Price": 199,
	//                 "Platform": "Tube",
	//                 "Password": "abcqwe",
	//                 "Tags": [
	//                         "JS",
	//                         "full"
	//                 ]
	//         },
	//         {
	//                 "Name": "Angular",
	//                 "Price": 99,
	//                 "Platform": "You cube Tube",
	//                 "Password": "123qwe",
	//                 "Tags": null
	//         }
	// ]

	// --------------------------------- * PROPERLY formatted * ------------------------------------
	// 	[
	//         {
	//                 "courseName": "ReactJs",
	//                 "Price": 299,
	//                 "websites": "YouTube",
	//                 "tags": [
	//                         "JS",
	//                         "web"
	//                 ]
	//         },
	//         {
	//                 "courseName": "MERN",
	//                 "Price": 199,
	//                 "websites": "Tube",
	//                 "tags": [
	//                         "JS",
	//                         "full"
	//                 ]
	//         },
	//         {
	//                 "courseName": "Angular",
	//                 "Price": 99,
	//                 "websites": "You cube Tube"
	//         }
	// ]

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", response)
}

func DecodeJson() {

	jsondatabytesfromweb := []byte(`
		{
			"courseName": "ReactJs",
			"Price": 299,
			"websites": "YouTube",
			"tags": ["JS","web"]
		}`,
	)

	var mycourse course

	isValid := json.Valid(jsondatabytesfromweb)
	if isValid {
		fmt.Println("Valid JSON")
		json.Unmarshal(jsondatabytesfromweb, &mycourse)

		fmt.Printf("%#v \n", mycourse) // it included hidden password
		// main.course{Name:"ReactJs", Price:299, Platform:"YouTube", Password:"", Tags:[]string{"JS", "web"}}

		fmt.Println(mycourse) // {ReactJs 299 YouTube  [JS web]}

	} else {
		fmt.Println("INVALID JSON")

	}

	// IN some cases where we don't follow struct parsing but instead individual key values for usage

	var myWebData map[string]interface{}
	// key is string, but when we are not aware whats the value is we can use interface

	json.Unmarshal(jsondatabytesfromweb, &myWebData)

	fmt.Printf("%#v \n", myWebData)
	// map[string]interface {}{"Price":299, "courseName":"ReactJs", "tags":[]interface {}{"JS", "web"}, "websites":"YouTube"}

	for k, v := range myWebData {
		fmt.Printf("The key `%v` is having value `%v` of Type: `%T` \n", k, v, v)
		// The key `courseName` is having value `ReactJs` of Type: `string`
		// The key `Price` is having value `299` of Type: `float64`
		// The key `websites` is having value `YouTube` of Type: `string`
		// The key `tags` is having value `[JS web]` of Type: `[]interface {}`
	}

}
