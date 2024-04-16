package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := []byte("\"date: " + ct.Format("02.01.2006") + ", time: " + ct.Format("15:04") + "\"")
	return formatted, nil
}

func getOk(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC()
	currentTime = currentTime.Add(3 * time.Hour)

	jsonResponse, err := json.Marshal(CustomTime{currentTime})
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
