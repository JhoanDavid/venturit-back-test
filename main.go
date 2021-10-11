package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	YOUR_API_KEY := getEnvVariable("OMDB_API_KEY")
	fmt.Println(YOUR_API_KEY)
	/*
		api := gomdb.Init(YOUR_API_KEY)
		query := &gomdb.QueryData{Title: "Macbeth", SearchType: gomdb.MovieSearch}
		res, err := api.Search(query)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.Search)

		query = &gomdb.QueryData{Title: "Macbeth", Year: "2015"}
		res2, err := api.MovieByTitle(query)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res2)

		res3, err := api.MovieByImdbID("tt2884018")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res3)*/
}

func getEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
