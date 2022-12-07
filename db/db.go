package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	dbCredential := model.CredentialDB{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "fitrisa02",
		DatabaseName: "Person",
		Port:         "5432",
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbCredential.Host, dbCredential.Username, dbCredential.Password, dbCredential.DatabaseName, dbCredential.Port)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
