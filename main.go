package main

import (
	"fmt"
	"net/http"
	"quoteapp/controller"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/routes"
)

func main() {
	db := db.NewDB(
		"192.168.0.101",
		"postgres",
		"password",
		"quoteapp",
		5432,
	)

	db.Migrate()
	conn := db.DB()

	model := model.NewQuoteModel(conn)
	controller := controller.NewQuoteController(model)

	router := routes.NewRoute(controller)

	fmt.Println("starting api server at http://localhost:8080")
	http.ListenAndServe(":8080", router.Run())
}
