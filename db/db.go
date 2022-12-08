package db

import (
	"ridwanHS/model"
	
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB()(*gorm.DB, error){
	dbCredential := model.CredentialDB{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "warva14",
		DatabaseName: "test",
		Port:         "5432",
	}
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbCredential.Host, dbCredential.Username, dbCredential.Password, dbCredential.DatabaseName, dbCredential.Port)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
