package controller

import (
	"encoding/json"
	"movies/helper"
	"movies/model"
	"movies/service"
	"net/http"
	"strconv"

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

func EditMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	err := service.EditMovie(movie)
	if err == nil {
		movieUpdated, err := service.GetMovieById(strconv.FormatInt(movie.Id, 10))
		if err == nil {
			helper.RespondWithSuccess(movieUpdated, w)
		} else {
			helper.RespondWithCustomError("Error, movie with id="+strconv.FormatInt(movie.Id, 10)+" doesn't exist", w)
		}
	} else {
		helper.RespondWithError(err, w)
	}
}

func GetFiltredMovies(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	if (keys["initial_released_year"] != nil && keys["final_released_year"] == nil) || (keys["initial_released_year"] == nil && keys["final_released_year"] != nil) {
		helper.RespondWithCustomError("Error, please verify initial_released_year and final_released_year are sent", w)
		return
	}

	filters := []string{}

	//released_year filter
	if keys["released_year"] != nil {
		filters = append(filters, "released_year="+keys["released_year"][0])
	} else if keys["initial_released_year"] != nil && keys["final_released_year"] != nil {

		if keys["initial_released_year"][0] > keys["final_released_year"][0] {
			helper.RespondWithCustomError("Error, please initial_released_year should be less than final_released_year", w)
		}
		filters = append(filters, "released_year>="+keys["initial_released_year"][0]+" AND released_year<="+keys["final_released_year"][0])
	}

	//rating filter
	if keys["rating"] != nil && keys["rating_especification"] == nil {
		filters = append(filters, "rating="+keys["rating"][0])
	} else if keys["rating"] != nil && keys["rating_especification"] != nil {
		switch keys["rating_especification"][0] {
		case "lower":
			filters = append(filters, "rating <= "+keys["rating"][0])
			break
		case "higher":
			filters = append(filters, "rating >= "+keys["rating"][0])
			break
		default:
			helper.RespondWithCustomError("Error, please send lower or higher for rating_especification", w)
		}
	}

	if keys["genres"] != nil {
		for i := range keys["genres"] {
			filters = append(filters, "genres like '%"+keys["genres"][i]+"%'")
		}
	}

	movieList, err := service.GetFiltredMovies(filters)
	if err == nil {
		helper.RespondWithSuccess(movieList, w)
	} else {
		helper.RespondWithError(err, w)
	}
}
