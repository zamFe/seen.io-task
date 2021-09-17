//TODO: Use dbresolver to manage multiple databases

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model

	Email string

	PhoneNumber string
}

type EventLog struct {
	gorm.Model

	LandingPageHits int

	VideoPlays int

	UserID int
}

var db *gorm.DB

var err error

var (
	users = []User{

		{Email: "james@email.com", PhoneNumber: "00000000"},

		{Email: "chris@email.com", PhoneNumber: "11111111"},

		{Email: "jenny@email.com", PhoneNumber: "22222222"},
	}

	eventLogs = []EventLog{

		{LandingPageHits: 1, VideoPlays: 1, UserID: 1},

		{LandingPageHits: 24, VideoPlays: 36, UserID: 2},

		{LandingPageHits: 63, VideoPlays: 125, UserID: 3},
	}
)

func main() {

	router := mux.NewRouter()
	dbSource := fmt.Sprintf("host=localhost port=5432 user=%s dbname=communications sslmode=disable password=%s", DbUser, DbPassword)
	db, err = gorm.Open("postgres", dbSource)

	if err != nil {

		panic(err)

	}

	defer db.Close()

	db.AutoMigrate(&User{})

	db.AutoMigrate(&EventLog{})

	for index := range users {

		db.Create(&users[index])

	}

	for index := range eventLogs {

		db.Create(&eventLogs[index])

	}

	router.HandleFunc("/eventLogs", GetEventLogs).Methods("GET")

	router.HandleFunc("/eventLogs/{id}", GetEventLog).Methods("GET")

	router.HandleFunc("/users", GetUsers).Methods("GET")

	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	router.HandleFunc("/eventLogs/{id}", DeleteEventLog).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func GetEventLogs(w http.ResponseWriter, r *http.Request) {

	var eventLogs []EventLog

	db.Find(&eventLogs)

	json.NewEncoder(w).Encode(&eventLogs)

}

func GetEventLog(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var eventLog EventLog

	db.First(&eventLog, params["id"])

	json.NewEncoder(w).Encode(&eventLog)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []User

	db.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var user User

	db.First(&user, params["id"])

	json.NewEncoder(w).Encode(&user)

}

func DeleteEventLog(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var eventLog EventLog

	db.First(&eventLog, params["id"])

	db.Delete(&eventLog)

	var eventLogs []EventLog

	db.Find(&eventLog)

	json.NewEncoder(w).Encode(&eventLogs)

}
