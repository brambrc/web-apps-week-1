package routes

import (
	"net/http"
	"quoteapp/controller"
)

type Route struct {
	mux   *http.ServeMux
	quote *controller.Quote
}

func NewRoute(q *controller.Quote) *Route {
	route := &Route{
		mux:   http.NewServeMux(),
		quote: q,
	}

	route.routes()
	return route
}

func (r *Route) Run() *http.ServeMux {
	return r.mux
}

func (r *Route) routes() {
	r.mux.Handle("/fetch", r.get(http.HandlerFunc(r.quote.Fetch)))
	r.mux.Handle("/select", r.get(http.HandlerFunc(r.quote.Show)))
	r.mux.Handle("/count", r.get(http.HandlerFunc(r.quote.Count)))
	r.mux.Handle("/add", r.post(http.HandlerFunc(r.quote.Store)))
}
