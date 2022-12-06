package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"webapp/model"
)

func (api *API) FetchQuote(w http.ResponseWriter, r *http.Request) {
	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var client = &http.Client{}

	// http.NewRequest() digunakan untuk membuat request baru
	hitTheApis, err := http.NewRequest("GET", "https://animechan.vercel.app/api/quotes", nil)
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
	var quote []model.Quotes
	var err2 = json.Unmarshal(responseData, &quote)
	if err2 != nil {
		fmt.Println(err2)
	}

	// cek error dan send response
	er := api.quoteRepo.FetchQuote(quote)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Fetch Data Gagal"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fetch Data Berhasil"))

}

func (api *API) SelectQuote(w http.ResponseWriter, r *http.Request) {
	resp, err := api.quoteRepo.SelectQuote()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Select Data Gagal"))
	}
	temporary := model.Get{
		Anime:     resp[0].Anime,
		Character: resp[0].Character,
		Quote:     resp[0].Quote,
	}
	data, er := json.Marshal(temporary)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Select Data Gagal Banget"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (api *API) Count(w http.ResponseWriter, r *http.Request) {
	resp, err := api.quoteRepo.Count()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Count Data Gagal"))
	}
	temporary := model.Count{
		Total: resp,
	}
	data, er := json.Marshal(temporary)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Count Data Gagal Banget"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (api *API) Add(w http.ResponseWriter, r *http.Request) {
	var quote model.Quotes
	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Data not added"))

	}
	er := api.quoteRepo.Add(quote)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Data not added"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully Add Data"))
}
