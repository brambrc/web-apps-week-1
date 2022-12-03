package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webApi/app/model"
)

func (api *API) FetchQuote(w http.ResponseWriter, r *http.Request) {
	var client = &http.Client{}

	hitTheApis, err := http.NewRequest("GET", "https://animechan.vercel.app/api/quotes", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(hitTheApis)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// Kita bisa membaca response body menggunakan package ioutil.
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var quotes []model.Quotes
	var err2 = json.Unmarshal(responseData, &quotes)
	if err2 != nil {
		fmt.Println(err2)
	}

	er := api.quoteRepo.FetchQuote(quotes)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Fetch Data Gagal"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fetch Data Berhasil"))

	//disini kita kirim response ke client

	// w.Write(wrap)
}
func (api *API) GetQuote(w http.ResponseWriter, r *http.Request) {
	resp, err := api.quoteRepo.GetQuote()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Get Data Gagal"))
	}
	temp := model.Get{
		Anime: resp.Anime,
		Char:  resp.Char,
		Quote: resp.Quote,
	}
	data, er := json.Marshal(temp)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Decode Gagal"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func (api *API) CountQuote(w http.ResponseWriter, r *http.Request) {
	count, err := api.quoteRepo.CountQuote()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Count Data Gagal"))
	}

	temp := model.Count{
		Total: count,
	}

	data, er := json.Marshal(temp)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Decode Gagal"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
func (api *API) AddQuote(w http.ResponseWriter, r *http.Request) {
	var data model.Quotes

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Decode Gagal"))

	}
	er := api.quoteRepo.AddQuote(data)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Add Data Gagal"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Add Data Berhasil"))
}
