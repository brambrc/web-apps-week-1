package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Quotes struct {
	Quote     string `json:"q"`
	Speaker   string `json:"a"`
	QuoteHTML string `json:"h"`
}

type AnimeQuotes struct {
	Anime string `json:"anime"`
	Char  string `json:"character"`
	Quote string `json:"quote"`
}

func getQuotes(w http.ResponseWriter, r *http.Request) {

	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var client = &http.Client{}

	// http.NewRequest() digunakan untuk membuat request baru
	hitTheApis, err := http.NewRequest("GET", "https://zenquotes.io/api/random", nil)
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
	var quote []Quotes
	var err2 = json.Unmarshal(responseData, &quote)
	if err2 != nil {
		fmt.Println(err2)
	}

	//disini kita convert struct Quotes ke json
	quo := &Quotes{
		Quote:     quote[0].Quote,
		Speaker:   quote[0].Speaker,
		QuoteHTML: quote[0].QuoteHTML,
	}

	wrap, err := json.Marshal(quo)

	if err != nil {
		fmt.Println(err)
	}

	//disini kita kirim response ke client
	w.Header().Set("Content-Type", "application/json")
	w.Write(wrap)

}
func getQuotesFromAnime(w http.ResponseWriter, r *http.Request) {
	//silahkan bikin hall yang sama dengan endpoint /getQuotes tapi dengan endpoint dibawah ini
	//https://animechan.vercel.app/api/random
	// dokumentasi https://animechan.vercel.app/docs
	// pilih api yang bisa mendapatkan 10 quotes random sekaligus, dan tampilkan semuanya

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
	var quote []AnimeQuotes
	var err2 = json.Unmarshal(responseData, &quote)
	if err2 != nil {
		fmt.Println(err2)
	}

	//disini kita convert struct Quotes ke json

	data := make([]AnimeQuotes, 0)
	for i := 0; i < len(quote); i++ {
		quo := &AnimeQuotes{
			Anime: quote[i].Anime,
			Char:  quote[i].Char,
			Quote: quote[i].Quote,
		}
		_, err := json.Marshal(quo)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, *quo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func main() {
	http.HandleFunc("/getQuotes", getQuotes)
	http.HandleFunc("/getQuotesFromAnime", getQuotesFromAnime)

	log.Println("Listening...at port 3333")
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s \n", err)
		return
	}
}
