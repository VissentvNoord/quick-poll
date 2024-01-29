package poll

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var polls []Poll

func CreatePoll(w http.ResponseWriter, r *http.Request) {
	var requestBody PollRequest

	err1 := json.NewDecoder(r.Body).Decode(&requestBody)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(requestBody)

	var newPoll Poll = *NewPoll(&requestBody)

	polls = append(polls, newPoll)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPoll)

	logSessions()
}

func getPollByID(id string) (Poll, error) {
	for _, poll := range polls {
		if poll.ID == id {
			return poll, nil
		}
	}

	return Poll{}, fmt.Errorf("no poll found")
}

func GetPoll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pollID := params["id"]
	poll, err := getPollByID(pollID)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(poll)
}

func UserSubmission(w http.ResponseWriter, r *http.Request) {
	var requestBody PollSubmission

	err1 := json.NewDecoder(r.Body).Decode(&requestBody)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	pollID := params["id"]
	poll, err := getPollByID(pollID)

	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	SubmitOption(requestBody.Submission, &poll)
}

func generatePollID() string {
	const idLength = 10
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.NewSource(time.Now().UnixNano())

	b := make([]byte, idLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func logSessions() {
	fmt.Printf("Current active polls: %v\n", len(polls))
	for _, poll := range polls {
		fmt.Println(poll)
	}
}
