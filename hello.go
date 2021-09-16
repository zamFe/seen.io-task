package main

import (
	"encoding/json"

	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"

	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Client struct {
	gorm.Model

	email string

	phoneNumber string

	EventLogs []EventLog
}

type EventLog struct {
	gorm.Model

	LandingPageHits int

	VideoPlays int
	ClientID   int
}

var db *gorm.DB

var err error

var (
	clients = []Client{

		{email: "james@email.com", phoneNumber: "00000000"},

		{email: "chris@email.com", phoneNumber: "11111111"},

		{email: "jenny@email.com", phoneNumber: "22222222"},
	}

	eventLogs = []EventLog{

		{LandingPageHits: 1, VideoPlays: 1, ClientID: 1},

		{LandingPageHits: 24, VideoPlays: 36, ClientID: 2},

		{LandingPageHits: 7, VideoPlays: 15, ClientID: 2},

		{LandingPageHits: 63, VideoPlays: 125, ClientID: 3},
	}
)

func main() {

	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=Mokka987!")

	if err != nil {

		panic(err)

	}

	defer db.Close()

	db.AutoMigrate(&Client{})

	db.AutoMigrate(&EventLog{})

	for index := range clients {

		db.Create(&clients[index])

	}

	for index := range eventLogs {

		db.Create(&eventLogs[index])

	}

	router.HandleFunc("/cars", GetCars).Methods("GET")

	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")

	router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")

	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func GetCars(w http.ResponseWriter, r *http.Request) {

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	json.NewEncoder(w).Encode(&car)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var driver Driver

	var cars []Car

	db.First(&driver, params["id"])

	db.Model(&driver).Related(&cars)

	driver.Cars = cars

	json.NewEncoder(w).Encode(&driver)

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	db.Delete(&car)

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}
