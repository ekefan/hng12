package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Resp struct {
	Email           string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GithubURL       string `json:"github_url"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Resp{
			Email:           "eebenezer949@gmail.com",
			CurrentDatetime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			GithubURL:       "https://github.com/ekefan/hng12",
		}
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		slog.Error("Server returned a non nil error", "details", err)
	}
}
