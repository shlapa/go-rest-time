package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

func getOk(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC()
	currentTime = currentTime.Add(3 * time.Hour)

	timeResponse := TimeResponse{
		Date: currentTime.Format("02.01.2006"),
		Time: currentTime.Format("15:04"),
	}

	jsonResponse, err := json.Marshal(timeResponse)
	if err != nil {
		http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/", getOk)
	http.HandleFunc("/time", getTime)
	http.ListenAndServe(":8080", nil)
}
