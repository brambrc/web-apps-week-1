package model

import "gorm.io/gorm"

type CredentialDB struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

type QuotesAnime struct {
	ID        int    `gorm:"primaryKey"`
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
	gorm.Model
}

type Respon struct {
	Message string
}
