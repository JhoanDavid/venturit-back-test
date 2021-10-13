package main

import (
	"log"
	"movies/db"
	"movies/routes"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	con, err := db.GetDB()

	if err != nil {
		log.Println("error with database " + err.Error())
	} else {
		err = con.Ping()
		if err != nil {
			log.Println("error making conection to DB, error: " + err.Error())
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
