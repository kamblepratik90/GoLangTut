package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models for course - files
type Course struct {
	CourseId    string  `json:id`
	CourseName  string  `json:name`
	CoursePrice int     `json:price`
	Author      *Author `json:author`
}

type Author struct {
	AuthorName  string `json:name`
	AuthorEmail string `json:email`
}

// fake DB
var courses []Course

// middleware / helpers - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Go lang build API ")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// seed courses
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJs", CoursePrice: 111,
		Author: &Author{AuthorName: "Pratik", AuthorEmail: "a@b.com"}})

	courses = append(courses, Course{CourseId: "2", CourseName: "Angular", CoursePrice: 222,
		Author: &Author{AuthorName: "Patric", AuthorEmail: "x@y.com"}})

	// routes
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - files

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is the homepage ... root... API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses...")
	// Headers
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")

	// Header
	w.Header().Add("Content-Type", "application/json")

	// req params
	params := mux.Vars(r)

	fmt.Printf("req params are : %v", params)

	// find
	for _, course := range courses {
		// conditional check for requested id
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found for given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Add("Content-Type", "application/json")

	// validations
	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what if : body {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside JSON {}")
	}

	// generate uniq id -> convert -> string
	// append course in to courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Add("Content-Type", "application/json")

	// 1st - grab id from req
	params := mux.Vars(r)

	// loop, find, remove, add again
	for index, course := range courses {
		// find
		if course.CourseId == params["id"] {
			// remove
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]

			// add again
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found for update...")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Add("Content-Type", "application/json")

	// 1st - grab id from req
	params := mux.Vars(r)

	// loop, find, remove
	for index, course := range courses {
		// find
		if course.CourseId == params["id"] {
			// remove
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]

			json.NewEncoder(w).Encode("course deleted successfully.")
			return
		}
	}

	json.NewEncoder(w).Encode("No course found for delete...")
	return
}
