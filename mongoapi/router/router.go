package router

import (
	"github.com/gorilla/mux"
	"github.com/priyankappm99/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Netflix-related routes
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	// Product-related routes
	router.HandleFunc("/api/products", controller.GetMyAllProducts).Methods("GET")
	router.HandleFunc("/api/product/{id}", controller.GetProduct).Methods("GET")
	router.HandleFunc("/api/product", controller.CreateProduct).Methods("POST")

	return router
}
