package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("this Go lang modules")
	greeter()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", serveHome).Methods("GET")

	// Bind to a port and pass our router in
	// log.Fatal(http.ListenAndServe(":3000", r))
	log.Fatal(http.ListenAndServe("localhost:3000", r))

	// visit -> localhost:3000
}

func greeter() {
	fmt.Println("Hello from Go lang modules...")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is home route... WELCOME</h1>"))
}
