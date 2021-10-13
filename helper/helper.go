package helper

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RespondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func RespondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func RespondWithCustomError(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(data)
}

func ArrayGenrestoString(arrayGenres []string) string {
	genres := ""
	for i := range arrayGenres {
		if i != 0 {
			genres += ", "
		}
		genres += arrayGenres[i]
	}
	return genres
}

func GenresStringtoArray(genres string) []string {
	arrayGenres := strings.Split(genres, ",")
	for i := range arrayGenres {
		arrayGenres[i] = strings.TrimSpace(arrayGenres[i])
	}
	return arrayGenres
}
