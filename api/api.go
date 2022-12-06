package api

import (
	"fmt"
	"net/http"
	repo "webapp/repository"
)

type API struct {
	quotesRepo repo.QuoteRepository
	mux        *http.ServeMux
}

func NewAPI(quotesRepo repo.QuoteRepository) API {
	mux := http.NewServeMux()
	api := API{
		quotesRepo,
		mux,
	}

	mux.Handle("/fetch", http.HandlerFunc(api.FetchQuote))
	mux.Handle("/select", http.HandlerFunc(api.SelectQuote))
	mux.Handle("/count", http.HandlerFunc(api.Count))
	mux.Handle("/add", http.HandlerFunc(api.Add))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
