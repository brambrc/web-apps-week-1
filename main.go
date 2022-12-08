package main

import (
	"fmt"
	
	"ridwanHS/db"
	"ridwanHS/model"
	"ridwanHS/repository"
	"log"
	"net/http"
)

func main() {
	connect, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	connect.AutoMigrate(&model.QuotesAnime{})

	k := repository.NewQuoteRepository(connect)

	http.HandleFunc("/insert-quotes-from-api", k.InsertQuotesFromAPI)
	http.HandleFunc("/get-quotes-from-api", k.GetQuotesFromDb)
	http.HandleFunc("/get-total-quotes", k.GetTotalQuotes)
	http.HandleFunc("/insert-quote", k.InsertQuotes)

	log.Println("Listening...at port 3333")
	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s \n", err)
		return
	}

}
