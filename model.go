package main

import "time"

type User struct {
	SlackName   string    `json:"slack_name"`
	CurrentDay  string    `json:"current_day"`
	CurrentTime time.Time `json:"utc_time"`
	Track       string    `json:"track"`
	GitFileUrl  string    `json:"github_file_url"`
	GitRepoUrl  string    `json:"github_repo_url"`
	StatusCode  int       `json:"status_code"`
}
