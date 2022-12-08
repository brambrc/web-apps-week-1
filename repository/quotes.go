package repository

import (
	"anime_quotes/model"
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

func (q *QuoteRepository) InsertQuotesFromAPI(w http.ResponseWriter, r *http.Request) {

	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var client = &http.Client{}

	// http.NewRequest() digunakan untuk membuat request baru
	hitTheApis, err := http.NewRequest("GET", "https://animechan.vercel.app/api/random", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(hitTheApis)
	if err != nil {
		panic(err)
	}

	// Cetak status code request
	fmt.Println("Status: ", resp.Status)

	// Kita bisa membaca response body menggunakan package ioutil.
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//disini response bodynya kita unmarshall dan convert ke struct Quotes
	var quote model.QuotesAnime
	err = json.Unmarshal(responseData, &quote)
	if err != nil {
		fmt.Println(err)
	}

	// insert to database
	err = q.db.Create(&quote).Error

	// jika insert data gagal
	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataInJson)
		return
	}

	// jika insert data berhasil
	msg := &model.ResponseMessege{
		Msg: "success insert quote to db",
	}

	dataInJson, err2 := json.Marshal(msg)
	if err2 != nil {
		panic(err2)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataInJson)
}

func (q *QuoteRepository) GetQuotesFromDb(w http.ResponseWriter, r *http.Request) {
	var quote model.QuotesAnime
	var totalQuotes int
	err := q.db.Raw("SELECT COUNT(id) FROM quotes_animes").Scan(&totalQuotes).Error

	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(int(totalQuotes)-1) + 1
	err = q.db.Raw("SELECT * FROM quotes_animes WHERE id = ?", randomNumber).Scan(&quote).Error

	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataInJson)
		return
	}

	dataInJson, err2 := json.Marshal(quote)
	if err2 != nil {
		panic(err2)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataInJson)
}

func (q *QuoteRepository) GetTotalQuotes(w http.ResponseWriter, r *http.Request) {
	var totalQuotes int
	err := q.db.Raw("SELECT COUNT(id) FROM quotes_animes").Scan(&totalQuotes).Error

	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataInJson)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	succmsg := fmt.Sprintf("Total Quotes in Database: %d", totalQuotes)
	w.Write([]byte(succmsg))
}

func (q *QuoteRepository) InsertQuotes(w http.ResponseWriter, r *http.Request) {
	var quote model.QuotesAnime

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert to database
	err = q.db.Create(&quote).Error

	// jika insert data gagal
	if err != nil {
		msg := &model.ResponseMessege{
			Msg: "found error",
		}

		dataInJson, err2 := json.Marshal(msg)
		if err2 != nil {
			panic(err2)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataInJson)
		return
	}

	// jika insert data berhasil
	msg := &model.ResponseMessege{
		Msg: "success insert quote to db",
	}

	dataInJson, err2 := json.Marshal(msg)
	if err2 != nil {
		panic(err2)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataInJson)
}
