package model

import (
	"gorm.io/gorm"
)

type CredentialDB struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
	// Schema       string
}

type AnimeQuotes struct {
	ID        int    `gorm:"primaryKey"`
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
	gorm.Model
}

type ResponseMessage struct {
	Msg string
}
