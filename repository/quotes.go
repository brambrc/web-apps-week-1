package repository

import (
	"WepApp/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type QuotesRepo struct {
	db *gorm.DB
}

func NewQuotesRepo(db *gorm.DB) *QuotesRepo {
	return &QuotesRepo{db}
}

func (q *QuotesRepo) InsertQuotesFromAPI(w http.ResponseWriter, r *http.Request) {

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

	// Membaca response body menggunakan package ioutil.
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Response bodynya kita unmarshall dan convert ke struct Quotes
	var quote model.AnimeQuotes
	var err2 = json.Unmarshal(responseData, &quote)
	if err2 != nil {
		fmt.Println(err2)
	}

	// add to database
	err = q.db.Create(&quote).Error

	if err != nil {
		// give error response
		msg := model.ResponseMessage{
			Msg: "Error to insert quote from API",
		}

		dataInJson, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dataInJson)
	}

	// give succes response
	msg := model.ResponseMessage{
		Msg: "success insert quote from API",
	}

	dataInJson, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataInJson)
}

func (q *QuotesRepo) ShowRandomQuote(w http.ResponseWriter, r *http.Request) {
	var listQuotes model.AnimeQuotes
	var total int64

	// get maks random from count quotes
	err := q.db.Table("anime_quotes").Count(&total).Error
	if err != nil {
		msg := model.ResponseMessage{
			Msg: "Found Error",
		}

		dataInJson, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dataInJson)
	}

	// random
	random := rand.Intn(int(total)-1) + 1

	// get random quote
	err = q.db.Table("anime_quotes").Where("id = ?", random).Scan(&listQuotes).Error
	if err != nil {
		// give error response
		msg := model.ResponseMessage{
			Msg: "Error to get quote",
		}

		dataInJson, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dataInJson)
	}

	// give success response
	dataInJson, err := json.Marshal(listQuotes)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataInJson)

}

func (q *QuotesRepo) CountQuotes(w http.ResponseWriter, r *http.Request) {
	var total int64

	// count quote with count function
	err := q.db.Table("anime_quotes").Count(&total).Error
	if err != nil {
		// give error response
		msg := model.ResponseMessage{
			Msg: "Error to count quotes",
		}

		dataInJson, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dataInJson)
	}

	// give success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Total quotes in database amounted to %d quotes ", total)))

}

func (q *QuotesRepo) InsertQuotes(w http.ResponseWriter, r *http.Request) {
	var quote model.AnimeQuotes

	// read response body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// unmarshall response to struct
	err = json.Unmarshal(b, &quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert to database
	err = q.db.Create(&quote).Error

	if err != nil {
		// give error response
		msg := &model.ResponseMessage{
			Msg: "error failed to insert data",
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

	// give success response
	msg := &model.ResponseMessage{
		Msg: "success insert quote to DB",
	}

	dataInJson, err2 := json.Marshal(msg)
	if err2 != nil {
		panic(err2)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataInJson)
}
