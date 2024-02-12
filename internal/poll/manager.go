package poll

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var polls []Poll

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Assuming you're receiving an integer
		if messageType == websocket.TextMessage {
			// Convert the received data (string) to an integer
			data, err := strconv.Atoi(string(p))
			if err != nil {
				fmt.Println("Failed to convert to integer:", err)
				return
			}
			// Handle the integer received from the client
			fmt.Println(data)
		}
	}
}

func DebugRequestBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	fmt.Println("Raw Request:", string(body))
}

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

	pollOption, err := SubmitOption(requestBody.Submission, &poll)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pollOption)
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
