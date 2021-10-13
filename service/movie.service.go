package service

import (
	"movies/db"
	"movies/helper"
	"movies/model"
	"os"

	"github.com/eefret/gomdb"
)

func CreateMovie(movie model.Movie) error {
	bd, err := db.GetDB()
	if err != nil {
		return err
	}
	genres := helper.ArrayGenrestoString(movie.Genres)
	_, err = bd.Exec("INSERT INTO movies (title, released_year, rating, genres) VALUES (?, ?, ?, ?)", movie.Title, movie.Released_year, movie.Rating, genres)
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
		movie.Genres = helper.GenresStringtoArray(genres)
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
	movie.Genres = helper.GenresStringtoArray(genres)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func GetMovieByTitle(title string) (model.Movie, error) {
	var movie model.Movie
	var genres string
	bd, err := db.GetDB()
	if err != nil {
		return movie, err
	}

	rows := bd.QueryRow("SELECT id, title, released_year, rating, genres FROM movies WHERE title=?", title)
	err = rows.Scan(&movie.Id, &movie.Title, &movie.Released_year, &movie.Rating, &genres)
	movie.Genres = helper.GenresStringtoArray(genres)
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
		movie.Genres = helper.GenresStringtoArray(genres)
		if err != nil {
			return moviesList, err
		}
		moviesList = append(moviesList, movie)
	}
	return moviesList, nil
}

func EditMovie(movie model.Movie) error {
	bd, err := db.GetDB()
	if err != nil {
		return err
	}
	genres := helper.ArrayGenrestoString(movie.Genres)
	_, err = bd.Exec("UPDATE movies SET rating=?, genres=? WHERE id=?", movie.Rating, genres, movie.Id)
	return err
}

func GetMovieOMDBByTitle(title string) (*gomdb.MovieResult, error) {
	OMDB_API_KEY := os.Getenv("OMDB_API_KEY")

	api := gomdb.Init(OMDB_API_KEY)
	query := &gomdb.QueryData{Title: title, SearchType: gomdb.MovieSearch}
	res, err := api.MovieByTitle(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
