package main

import (
	"anime_quotes/db"
	"anime_quotes/model"
	"anime_quotes/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	conn, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.QuotesAnime{})

	q := repository.NewQuoteRepository(conn)

	http.HandleFunc("/insert-quotes-from-api", q.InsertQuotesFromAPI)
	http.HandleFunc("/get-quotes-from-api", q.GetQuotesFromDb)
	http.HandleFunc("/get-total-quotes", q.GetTotalQuotes)
	http.HandleFunc("/insert-quote", q.InsertQuotes)

	log.Println("Listening...at port 3333")
	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s \n", err)
		return
	}

}
