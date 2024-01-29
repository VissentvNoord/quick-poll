package main

import (
	"net/http"
	"quick-poll/internal/poll"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/new", poll.CreatePoll).Methods("POST")
	r.HandleFunc("/{id}", poll.GetPoll).Methods("GET")
	r.HandleFunc("/{id}/submit", poll.UserSubmission).Methods("POST")

	http.Handle("/", r)

	port := ":8080"
	println("Server listening on port", port)
	http.ListenAndServe(port, nil)
}
