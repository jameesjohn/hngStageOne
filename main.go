package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type response struct {
	SlackName     string `json:"slack_name"`
	Track         string `json:"track"`
	CurrentDay    string `json:"current_day"`
	UtcTime       string `json:"utc_time"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func success(w http.ResponseWriter, data interface{}) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		fail(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
}

func fail(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := fmt.Fprint(w, "Unable to process request", data)
	if err != nil {
		log.Println("unhandled error:", err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received to", r.URL.Path)

	queries := r.URL.Query()
	slackName := queries.Get("slack_name")
	track := queries.Get("track")

	log.Printf("slackName: %v, track: %v\n", slackName, track)

	now := time.Now()
	responseValue := response{
		SlackName:     slackName,
		Track:         track,
		CurrentDay:    now.Format("Monday"),
		UtcTime:       now.UTC().Format(time.RFC3339),
		GithubFileUrl: "github.something",
		GithubRepoUrl: "https://github.com/jameesjohn/hngStageOne",
		StatusCode:    http.StatusOK,
	}

	success(w, responseValue)
}

func main() {
	http.HandleFunc("/api", handler)

	log.Println("Go server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
