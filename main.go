package main

import (
	api "webApi/app/controller"
	"webApi/app/model"
	repo "webApi/app/repository"
	"webApi/config"
)

func main() {
	db := config.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "asd123",
		DatabaseName: "pgadmin",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.Quotes{})

	quotesRepo := repo.NewQuoteRepo(conn)

	mainAPI := api.NewAPI(quotesRepo)
	mainAPI.Start()
}
