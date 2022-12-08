package repository

import (
	"webapp/model"

	"gorm.io/gorm"
)

type QuoteRepository struct {
	db *gorm.DB
}

func NewQuotesRepository(db *gorm.DB) QuoteRepository {
	return QuoteRepository{db}
}

func (u *QuoteRepository) FetchQuote(quote []model.Quotes) error {
	return u.db.Create(&quote).Error
}

func (u *QuoteRepository) SelectQuote() ([]model.Quotes, error) {
	quotes := []model.Quotes{}
	err := u.db.Raw("SELECT anime,character,quote FROM quotes ORDER BY RANDOM() LIMIT 1").Scan(&quotes).Error
	if err != nil {
		return []model.Quotes{}, err
	}
	return quotes, nil
}

func (u *QuoteRepository) Count() (int64, error) {
	var count int64
	err := u.db.Model(&model.Quotes{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *QuoteRepository) Add(quote model.Quotes) error {
	err := u.db.Create(&quote).Error
	if err != nil {
		return err
	}
	return nil
}
