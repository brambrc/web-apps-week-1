package controller

import (
	"fmt"
	"net/http"
	"os"

	repo "webApi/app/repository"
)

type API struct {
	quoteRepo repo.QuoteRepo
	mux       *http.ServeMux
}

func NewAPI(quoteRepo repo.QuoteRepo) API {
	mux := http.NewServeMux()
	api := API{
		quoteRepo,
		mux,
	}

	mux.Handle("/fetch", http.HandlerFunc(api.FetchQuote))
	mux.Handle("/get", http.HandlerFunc(api.GetQuote))
	mux.Handle("/count", http.HandlerFunc(api.CountQuote))
	mux.Handle("/add", http.HandlerFunc(api.AddQuote))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":"+port, api.Handler())
}
