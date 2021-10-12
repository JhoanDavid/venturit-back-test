package main

import (
	"fmt"
	"log"
	"movies/db"
	"movies/routes"
	"net/http"
	"os"
	"time"

	"github.com/eefret/gomdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	con, err := db.GetDB()

	if err != nil {
		fmt.Println("error with database " + err.Error())
	} else {
		err = con.Ping()
		if err != nil {
			fmt.Println("error making conection to DB, error: " + err.Error())
			return
		}
	}

	router := mux.NewRouter()
	routes.SetupRoutes(router)

	port := ":8000"

	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
}

func getMovieOMDBByTitle(title string) {
	OMDB_API_KEY := os.Getenv("OMDB_API_KEY")

	api := gomdb.Init(OMDB_API_KEY)
	query := &gomdb.QueryData{Title: title, SearchType: gomdb.MovieSearch}
	res, err := api.MovieByTitle(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
