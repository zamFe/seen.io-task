//TODO: Use dbresolver to manage multiple databases

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"seenio/dbContext/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

var err error

var (
	users = []model.User{

		{Email: "james@email.com", PhoneNumber: "00000000"},

		{Email: "chris@email.com", PhoneNumber: "11111111"},

		{Email: "jenny@email.com", PhoneNumber: "22222222"},
	}

	eventLogs = []model.EventLog{

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

	db.AutoMigrate(&model.User{})

	db.AutoMigrate(&model.EventLog{})

	for index := range users {

		db.Create(&users[index])

	}

	for index := range eventLogs {

		db.Create(&eventLogs[index])

	}

	router.HandleFunc("/eventlogs", GetEventLogs).Methods("GET")

	router.HandleFunc("/eventlogs/{id}", GetEventLog).Methods("GET")

	router.HandleFunc("/users", GetUsers).Methods("GET")

	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	router.HandleFunc("/eventlogs/{id}", DeleteEventLog).Methods("DELETE")

	router.HandleFunc("/eventlogs/videoplays/{id}", UpdateVideoPlays).Methods("PATCH")

	router.HandleFunc("/eventlogs/landingpagehits/{id}", UpdateLandingPageHits).Methods("PATCH")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func GetEventLogs(w http.ResponseWriter, r *http.Request) {

	var eventLogs []model.EventLog

	db.Find(&eventLogs)

	json.NewEncoder(w).Encode(&eventLogs)

}

func GetEventLog(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var eventLog model.EventLog

	db.First(&eventLog, params["id"])

	json.NewEncoder(w).Encode(&eventLog)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []model.User

	db.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var user model.User

	db.First(&user, params["id"])

	json.NewEncoder(w).Encode(&user)

}

func UpdateLandingPageHits(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var eventLog model.EventLog

	json.NewDecoder(r.Body).Decode(&eventLog)

	db.Model(&eventLog).Where("user_id = ?", params["id"]).Update("landing_page_hits", eventLog.LandingPageHits)

	json.NewEncoder(w).Encode((&eventLog))
}

func UpdateVideoPlays(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var eventLog model.EventLog

	json.NewDecoder(r.Body).Decode(&eventLog)

	db.Model(&eventLog).Where("user_id = ?", params["id"]).Update("video_plays", eventLog.VideoPlays)

	json.NewEncoder(w).Encode((&eventLog))
}

func DeleteEventLog(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var eventLog model.EventLog

	db.First(&eventLog, params["id"])

	db.Delete(&eventLog)

	var eventLogs []model.EventLog

	db.Find(&eventLog)

	json.NewEncoder(w).Encode(&eventLogs)

}
