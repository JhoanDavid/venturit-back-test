package routes

import (
	"movies/controller"
	"movies/helper"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.RespondWithSuccess("Venturit-test Api rest", w)
	}).Methods(http.MethodGet)

	router.HandleFunc("/movies", controller.GetAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", controller.GetMovieById).Methods(http.MethodGet)

}
