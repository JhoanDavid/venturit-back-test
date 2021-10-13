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

	subRouter := router.PathPrefix("/movies").Subrouter()
	subRouter.HandleFunc("/all", controller.GetAllMovies).Methods(http.MethodGet)
	subRouter.HandleFunc("/filtered", controller.GetFiltredMovies).Methods(http.MethodGet)
	subRouter.HandleFunc("/edit", controller.EditMovie).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", controller.GetMovieById).Methods(http.MethodGet)

}
