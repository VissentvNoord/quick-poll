package poll

import (
	"fmt"
	"time"
)

type PollSubmission struct {
	Submission int `json:"submission"`
}

type PollRequest struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
}

type Poll struct {
	ID        string       `json:"id"`
	Title     string       `json:"title"`
	CreatedAt time.Time    `json:"created_at"`
	Duration  float32      `json:"duration"`
	Options   []PollOption `json:"options"`
}

type PollOption struct {
	Description string `json:"descr"`
	Submissions int    `json:"submissions"`
}

func NewPoll(req *PollRequest) *Poll {
	var pollOptions []PollOption

	for _, optionString := range req.Options {
		pollOption := PollOption{
			Description: optionString,
		}

		pollOptions = append(pollOptions, pollOption)
	}

	return &Poll{
		ID:        generatePollID(),
		Title:     req.Title,
		CreatedAt: time.Now(),
		Options:   pollOptions,
	}
}

func SubmitOption(option int, poll *Poll) {
	fmt.Printf("Length of options: %v\n", len(poll.Options))

	if option >= len(poll.Options) {
		fmt.Printf("Wrong submission, length of options too short\n")
		return
	}

	poll.Options[option].Submissions++
	fmt.Printf("Submissed option: %s\n", poll.Options[option].Description)

	fmt.Printf("Current poll standings: \n")

	for _, o := range poll.Options {
		fmt.Printf("%s: %v\n", o.Description, o.Submissions)
	}
}
