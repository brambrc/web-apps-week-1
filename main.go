package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas/repository"
)

func main() {
	con, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	db.Automigrate(&model.QuotesAnime{})

	quo := repository.NewQuote(con)

	http.HandleFunc("/insert-quotes-from-api", quo.InsertQuotesFromApi)
	http.HandleFunc("/get-quotes-from-api", quo.GetQuotesFromDB)
	http.HandleFunc("/get-total-quotes", quo.GetTotalQuotes)
	http.HandleFunc("/insert-quote", quo.InsertQuotes)

	log.Println("Listening...at port 3333")
	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s \n", err)
		return
	}
}
