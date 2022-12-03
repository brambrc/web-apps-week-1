package repository

import (
	"webApi/app/model"

	"gorm.io/gorm"
)

type QuoteRepo struct {
	db *gorm.DB
}

func NewQuoteRepo(db *gorm.DB) QuoteRepo {
	return QuoteRepo{db}
}

func (u *QuoteRepo) FetchQuote(quotes []model.Quotes) error {
	if err := u.db.Create(quotes).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *QuoteRepo) GetQuote() (model.Quotes, error) {
	result := model.Quotes{}

	if err := u.db.Raw("SELECT anime,char,quote FROM quotes ORDER BY RANDOM() LIMIT 1").Scan(&result).Error; err != nil {
		return model.Quotes{}, err
	}

	return result, nil // TODO: replace this
}

func (u *QuoteRepo) CountQuote() (int64, error) {
	var result int64
	if err := u.db.Model(&model.Quotes{}).Count(&result).Error; err != nil {
		return 0, err
	}

	return result, nil // TODO: replace this
}

func (u *QuoteRepo) AddQuote(quote model.Quotes) error {
	if err := u.db.Create(&quote).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
