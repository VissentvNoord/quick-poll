package main

import (
	"net/http"
	"quick-poll/internal/poll"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ws", poll.HandleWebSocket)
	r.HandleFunc("/new", poll.CreatePoll).Methods("POST")
	r.HandleFunc("/{id}", poll.GetPoll).Methods("GET")
	r.HandleFunc("/{id}/submit", poll.UserSubmission).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})

	corsRouter := handlers.CORS(headersOk, originsOk, methodsOk)

	http.Handle("/", corsRouter(r))

	port := ":8080"
	println("Server listening on port", port)
	http.ListenAndServe(port, nil)
}
