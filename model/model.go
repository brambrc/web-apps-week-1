package model

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type Quotes struct {
	ID        int    `gorm:"primaryKey"`
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}
