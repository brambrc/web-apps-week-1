package repository

import (
	"ridwanHS/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type QuoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) *QuoteRepository {
	return &QuoteRepository{
		db: db,
	}
}

func (j *QuoteRepository) InsertQuotesFromAPI(k http.ResponseWriter, r *http.Request) {


	var client = &http.Client{}

	
	hitTheApis, err := http.NewRequest("GET", "https://animechan.vercel.app/api/random", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(hitTheApis)
	if err != nil {
		panic(err)
	}

	
	fmt.Println("Status: ", resp.Status)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}


	var quote model.QuotesAnime
	err = json.Unmarshal(responseData, &quote)
	if err != nil {
		fmt.Println(err)
	}

	
	err = j.db.Create(&quote).Error

	
	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		k.Header().Set("Content-Type", "application/json")
		k.WriteHeader(500)
		k.Write(dataInJson)
		return
	}

	
	msg := &model.ResponseMessege{
		Msg: "success insert quote to db",
	}

	dataInJson, err2 := json.Marshal(msg)
	if err2 != nil {
		panic(err2)
	}
	k.Header().Set("Content-Type", "application/json")
	k.WriteHeader(200)
	k.Write(dataInJson)
}

func (j *QuoteRepository) GetQuotesFromDb(k http.ResponseWriter, r *http.Request) {
	var quote model.QuotesAnime
	var totalQuotes int
	err := j.db.Raw("SELECT COUNT(id) FROM quotes_animes").Scan(&totalQuotes).Error

	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(int(totalQuotes)-1) + 1
	err = j.db.Raw("SELECT * FROM quotes_animes WHERE id = ?", randomNumber).Scan(&quote).Error

	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		k.Header().Set("Content-Type", "application/json")
		k.WriteHeader(500)
		k.Write(dataInJson)
		return
	}

	dataInJson, err2 := json.Marshal(quote)
	if err2 != nil {
		panic(err2)
	}
	k.Header().Set("Content-Type", "application/json")
	k.WriteHeader(200)
	k.Write(dataInJson)
}

func (j *QuoteRepository) GetTotalQuotes(k http.ResponseWriter, r *http.Request) {
	var totalQuotes int
	err := j.db.Raw("SELECT COUNT(id) FROM quotes_animes").Scan(&totalQuotes).Error

	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		k.Header().Set("Content-Type", "application/json")
		k.WriteHeader(500)
		k.Write(dataInJson)
		return
	}

	k.Header().Set("Content-Type", "application/json")
	k.WriteHeader(200)
	succmsg := fmt.Sprintf("Total Quotes in Database: %d", totalQuotes)
	k.Write([]byte(succmsg))
}

func (j *QuoteRepository) InsertQuotes(k http.ResponseWriter, r *http.Request) {
	var quote model.QuotesAnime

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(k, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &quote)
	if err != nil {
		http.Error(k, err.Error(), http.StatusBadRequest)
		return
	}

	
	err = j.db.Create(&quote).Error

	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		k.Header().Set("Content-Type", "application/json")
		k.WriteHeader(500)
		k.Write(dataInJson)
		return
	}

	
	msg := &model.ResponseMessege{
		Msg: "success insert quote to db",
	}

	dataInJson, err2 := json.Marshal(msg)
	if err2 != nil {
		panic(err2)
	}
	k.Header().Set("Content-Type", "application/json")
	k.WriteHeader(200)
	k.Write(dataInJson)
}
