package repository

import (
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

func NewQuote(db *gorm.DB) *QuoteRepository {
	return &QuoteRepository{
		db: db,
	}
}

func (q *QuoteRepository) InsertQuotesFromApi(w http.ResponseWriter, r *http.Request) {
	//statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var Client = &http.Client{}

	//http.NewRequest() digunakan untuk membuat request baru
	hitApi, err := http.NewRequest("GET", "https://animechan.vercel.app/api/random", nil)
	if err != nil {
		panic(err)
	}

	respon, err := Client.Do(hitApi)
	if err != nil {
		panic(err)
	}

	//cetak status code request
	fmt.Println("Status: ", respon.Status)

	//kita bisa membaca response body menggunakan package ioutil
	responseData, err := ioutil.ReadAll(respon.Body)
	if err != nil {
		log.Fatal(err)
	}

	//disini response bodynya kita unmarshall dan convert ke struct Quotes
	var a model.QuotesAnime
	err = json.Unmarshal(responseData, &a)
	if err != nil {
		fmt.Println(err)
	}

	//insert to database
	err = q.db.Create(&q).Error

	//jika insert data gagal
	if err != nil {
		msg := &model.Respon{
			Message: "error",
		}

		dataJson, err1 := json.Marshal(msg)
		if err1 != nil {
			panic(err1)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataJson)
		return
	}

	//jika insert data berhasil
	msage := &model.Respon{
		Message: "Success insert quote to db",
	}

	dataJson, err2 := json.Marshal(msage)
	if err2 != nil {
		panic(err2)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataJson)
	return
}

func (q *QuoteRepository) GetQuotesFromDB(w http.ResponseWriter, r *http.Request) {
	var quotess model.QuotesAnime
	var TotalQuote int
	err := q.db.Raw("SELECT COUNT(id) FROM quotes_animes").Scan(&TotalQuote).Error

	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(int(TotalQuote)-1) + 1
	err = q.db.Raw("SELECT * FROM quotes_animes WHERE id = ?", randomNum).Scan(&quotess).Error

	if err != nil {
		msage := &model.Respon{
			Message: "error",
		}

		dataJson, err1 := json.Marshal(msage)
		if err1 != nil {
			panic(err1)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataJson)
		return
	}

	dataJson, err1 := json.Marshal(quotess)
	if err1 != nil {
		panic(err1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataJson)
	return

}

func (q *QuoteRepository) GetTotalQuotes(w http.ResponseWriter, r *http.Request) {
	var totalQ int
	err := q.db.Raw("SELECT COUNT(id) FROM quotes_animes").Scan(&totalQ).Error

	if err != nil {
		msage := &model.Respon{
			Message: "error",
		}

		dataJson, err1 := json.Marshal(msage)
		if err1 != nil {
			panic(err1)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataJson)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	succmsg := fmt.Sprintf("Total Quotes in Database: %d", totalQ)
	w.Write([]byte(succmsg))
}

func (q *QuoteRepository) InsertQuotes(w http.ResponseWriter, r *http.Request) {
	var quo model.QuotesAnime

	a, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.Unmarshal(a, &quo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//insert to database
	err = q.db.Create(&quo).Error

	//jika insert data gagal
	if err != nil {
		msage := &model.Respon{
			Message: "error",
		}

		dataJson, err1 := json.Marshal(msage)
		if err1 != nil {
			panic(err1)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dataJson)
		return
	}
	msage := &model.Respon{
		Message: "Success insert quote to db",
	}

	dataJson, err1 := json.Marshal(msage)
	if err1 != nil {
		panic(err1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataJson)
}
