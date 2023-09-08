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

/*func populatejson() error {
	user := &User{
		SlackName:   "ichthoth",
		CurrentDay:  time.Now().Day(),
		CurrentTime: time.Now().Local(),
		Track:       "backend",
		GitFileUrl:  "http://github.com/ichthoth/",
		GitRepoUrl:  "",
		StatusCode:  http.StatusOK,
	}

	_, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	return err
}*/

func GetSlack(w http.ResponseWriter, r *http.Request) {
	user := &User{
		SlackName:   "ichthoth",
		CurrentDay:  "Friday",
		CurrentTime: time.Now().UTC(),
		Track:       "backend",
		GitFileUrl:  "http://github.com/ichthoth/hngx-task1/blob/master/main.go",
		GitRepoUrl:  "http://github.com/ichthoth/hngx-task1",
		StatusCode:  http.StatusOK,
	}

	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "non get request", http.StatusNotFound)
	}

	name := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", u)
	fmt.Fprintf(w, "slack_name:%s\n", name)
	fmt.Fprintf(w, "track:%s\n", track)
}

func main() {
	http.HandleFunc("/home", GetSlack)
	fmt.Printf("startng server at 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
