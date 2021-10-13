package service

import (
	"movies/db"
	"movies/model"
	"strings"
)

func CreateMovie(movie model.Movie) error {
	bd, err := db.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO movies (title, released_year, rating, genres) VALUES (?, ?, ?, ?)", movie.Title, movie.Released_year, movie.Rating, movie.Genres)
	return err
}

func GetAllMovies() ([]model.Movie, error) {
	moviesList := []model.Movie{}
	bd, err := db.GetDB()
	if err != nil {
		return moviesList, err
	}

	rows, err := bd.Query("SELECT id, title, released_year, rating, genres FROM movies")
	if err != nil {
		return moviesList, err
	}

	for rows.Next() {
		var movie model.Movie
		var genres string
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Released_year, &movie.Rating, &genres)
		arrayGenres := strings.Split(genres, ",")
		for i := range arrayGenres {
			arrayGenres[i] = strings.TrimSpace(arrayGenres[i])
		}
		movie.Genres = arrayGenres
		if err != nil {
			return moviesList, err
		}
		moviesList = append(moviesList, movie)
	}
	return moviesList, nil
}

func GetMovieById(id string) (model.Movie, error) {
	var movie model.Movie
	var genres string
	bd, err := db.GetDB()
	if err != nil {
		return movie, err
	}

	rows := bd.QueryRow("SELECT id, title, released_year, rating, genres FROM movies WHERE id=?", id)
	err = rows.Scan(&movie.Id, &movie.Title, &movie.Released_year, &movie.Rating, &genres)
	arrayGenres := strings.Split(genres, ",")
	for i := range arrayGenres {
		arrayGenres[i] = strings.TrimSpace(arrayGenres[i])
	}
	movie.Genres = arrayGenres
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func GetFiltredMovies(filters []string) ([]model.Movie, error) {
	moviesList := []model.Movie{}
	bd, err := db.GetDB()
	if err != nil {
		return moviesList, err
	}
	query := "SELECT id, title, released_year, rating, genres FROM movies WHERE "

	for i := range filters {
		if i != 0 {
			query += " AND "
		}
		query += filters[i]
	}

	rows, err := bd.Query(query)
	if err != nil {
		return moviesList, err
	}

	for rows.Next() {
		var movie model.Movie
		var genres string
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Released_year, &movie.Rating, &genres)
		arrayGenres := strings.Split(genres, ",")
		for i := range arrayGenres {
			arrayGenres[i] = strings.TrimSpace(arrayGenres[i])
		}
		movie.Genres = arrayGenres
		if err != nil {
			return moviesList, err
		}
		moviesList = append(moviesList, movie)
	}
	return moviesList, nil
}

func EditMovie(id string) (model.Movie, error) {
	var movie model.Movie
	var genres string
	bd, err := db.GetDB()
	if err != nil {
		return movie, err
	}

	rows := bd.QueryRow("SELECT id, title, released_year, rating, genres FROM movies WHERE id=?", id)
	err = rows.Scan(&movie.Id, &movie.Title, &movie.Released_year, &movie.Rating, &genres)
	arrayGenres := strings.Split(genres, ",")
	for i := range arrayGenres {
		arrayGenres[i] = strings.TrimSpace(arrayGenres[i])
	}
	movie.Genres = arrayGenres
	if err != nil {
		return movie, err
	}
	return movie, nil
}
