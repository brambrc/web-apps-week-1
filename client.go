package main

// import (
// 	"WepApp/model"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	_ "github.com/lib/pq"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Quotes struct {
// 	Quote     string `json:"q"`
// 	Speaker   string `json:"a"`
// 	QuoteHTML string `json:"h"`
// }

// type AnimeRepo struct {
// 	db *gorm.DB
// }

// func NewAnimeRepo(db *gorm.DB) AnimeRepo {
// 	return AnimeRepo{db}
// }

// func connectToDB() (*gorm.DB, error) {
// 	dbCredential := model.CredentialDB{
// 		Host:         "localhost",
// 		Username:     "postgres",
// 		Password:     "HansPG001",
// 		DatabaseName: "ChallangeDatabase",
// 		Port:         "5432",
// 	}

// 	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbCredential.Host, dbCredential.Username, dbCredential.Password, dbCredential.DatabaseName, dbCredential.Port)
// 	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func getQuotes(w http.ResponseWriter, r *http.Request) {

// 	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
// 	var client = &http.Client{}

// 	// http.NewRequest() digunakan untuk membuat request baru
// 	hitTheApis, err := http.NewRequest("GET", "https://zenquotes.io/api/random", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	resp, err := client.Do(hitTheApis)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Cetak status code request
// 	fmt.Println("Status: ", resp.Status)

// 	// Kita bisa membaca response body menggunakan package ioutil.
// 	responseData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//disini response bodynya kita unmarshall dan convert ke struct Quotes
// 	var quote []Quotes
// 	var err2 = json.Unmarshal(responseData, &quote)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}

// 	//disini kita convert struct Quotes ke json
// 	quo := &Quotes{
// 		Quote:     quote[0].Quote,
// 		Speaker:   quote[0].Speaker,
// 		QuoteHTML: quote[0].QuoteHTML,
// 	}

// 	wrap, err := json.Marshal(quo)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	//disini kita kirim response ke client
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(wrap)

// }
// func getQuotesFromAnime(w http.ResponseWriter, r *http.Request) {
// 	//silahkan bikin hall yang sama dengan endpoint /getQuotes tapi dengan endpoint dibawah ini
// 	//dokumentasi https://animechan.vercel.app/docs
// 	//pilih api yang bisa mendapatkan 10 quotes random sekaligus, dan tampilkan semuanya
// 	//struktur functionnya sama dengan function getQuotes diatas, namun karena datanya lebih banyak, maka kita perlu menggunakan array of struct
// 	//dan juga perlu menggunakan looping untuk menampilkan semua quotes
// 	//hint: tambahkan struct baru yang sesuai dengan struktur json yang didapatkan dari api

// }

// func (a AnimeRepo) fetchQuotesFromAnime() {
// 	resp, err := http.Get("https://animechan.vercel.app/api/quotes")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	responseData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var AnimeQuotes []AnimeQuotes
// 	var err2 = json.Unmarshal(responseData, &AnimeQuotes)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}

// 	if err := a.db.Create(&AnimeQuotes).Error; err != nil {
// 		panic(err)
// 	}

// }

// func selectQuotesFromAnime(w http.ResponseWriter, r *http.Request) {

// }

// func countQuotesFromAnime(w http.ResponseWriter, r *http.Request) {

// }

// func addQuotesFromAnime(w http.ResponseWriter, r *http.Request) {

// }

// func main() {

// 	dbCredential := CredentialDB{
// 		Host:         "localhost",
// 		Username:     "postgres",
// 		Password:     "HansPG001",
// 		DatabaseName: "ChallangeDatabase",
// 		Port:         "5432",
// 	}

// 	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbCredential.Host, dbCredential.Username, dbCredential.Password, dbCredential.DatabaseName, dbCredential.Port)
// 	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}

// 	db.AutoMigrate(&AnimeQuotes{})

// 	http.HandleFunc("/getQuotes", getQuotes)
// 	http.HandleFunc("/getQuotesFromAnime", getQuotesFromAnime)
// 	http.HandleFunc("/fetch", fetchQuotesFromAnime)
// 	http.HandleFunc("/select", selectQuotesFromAnime)
// 	http.HandleFunc("/count", countQuotesFromAnime)
// 	http.HandleFunc("/add", addQuotesFromAnime)

// 	log.Println("Listening...at port 3333")
// 	err = http.ListenAndServe(":3333", nil)

// 	if err != nil {
// 		fmt.Printf("Error starting server: %s \n", err)
// 		return
// 	}
// }
