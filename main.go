package main

import (
	"WepApp/db"
	"WepApp/model"
	"WepApp/repository"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	conn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.AnimeQuotes{})
	repo := repository.NewQuotesRepo(conn)

	http.HandleFunc("/fetch", repo.InsertQuotesFromAPI)
	http.HandleFunc("/get", repo.ShowRandomQuote)
	http.HandleFunc("/count", repo.CountQuotes)
	http.HandleFunc("/add", repo.InsertQuotes)

	log.Println("Listening...at port 3333")
	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s \n", err)
		return
	}

	fmt.Print("success")

}
