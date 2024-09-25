package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kamblepratik90/mongoDbGoLang/router"
)

func main() {
	fmt.Println("Go lang Mongo DB...")
	r := router.Router()
	fmt.Println("Server is getting ready...")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening on port 4000....")
}
