package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	SlackName   string    `json:"slack_name"`
	CurrentDay  string    `json:"current_day"`
	CurrentTime time.Time `json:"utc_time"`
	Track       string    `json:"track"`
	GitFileUrl  string    `json:"github_file_url"`
	GitRepoUrl  string    `json:"github_repo_url"`
	StatusCode  int       `json:"status_code"`
}

func GetSlack(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "NON-GET REQUEST", http.StatusBadRequest)
	}

	query := r.URL.Query()

	name := query.Get("slack_name")
	track := query.Get("track")

	if name == "" || track == "" {
		http.Error(w, "cannot get query params", http.StatusBadRequest)
	}

	response := &User{
		SlackName:   "ichthoth",
		CurrentDay:  time.Now().Weekday().String(),
		CurrentTime: time.Now().UTC(),
		Track:       "backend",
		GitFileUrl:  "http://github.com/ichthoth/hngx-task1/blob/master/main.go",
		GitRepoUrl:  "http://github.com/ichthoth/hngx-task1",
		StatusCode:  http.StatusOK,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "couldnt parse json response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/api", GetSlack)
	fmt.Printf("starting server at 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
