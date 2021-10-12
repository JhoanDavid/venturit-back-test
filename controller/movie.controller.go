package controller

import (
	"movies/helper"
	"movies/service"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movieList, err := service.GetAllMovies()
	if err == nil {
		helper.RespondWithSuccess(movieList, w)
	} else {
		helper.RespondWithError(err, w)
	}
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	movie, err := service.GetMovieById(id)
	if err == nil {
		helper.RespondWithSuccess(movie, w)
	} else {
		helper.RespondWithError(err, w)
	}
}
