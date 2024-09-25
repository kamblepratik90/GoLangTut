package router

import (
	"github.com/gorilla/mux"
	"github.com/kamblepratik90/mongoDbGoLang/controller"
)

func Router() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/api", controller.ServeHome).Methods("GET")

	route.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	route.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	route.HandleFunc("/api/movie/{id}", controller.MarkMovieWatched).Methods("PUT")
	route.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	route.HandleFunc("/api/movies", controller.DeleteALLMovie).Methods("DELETE")

	return route
}
