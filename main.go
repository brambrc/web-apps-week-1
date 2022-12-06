package main

import (
	"webapp/api"
	"webapp/db"
	"webapp/model"
	repo "webapp/repository"
)

func main() {
	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "ormcp1",
		Port:         2502,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.Quotes{})

	quoteRepo := repo.NewQuotesRepository(conn)

	mainAPI := api.NewAPI(quoteRepo)
	mainAPI.Start()
}
