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

func logPollStandings(poll *Poll) {
	fmt.Printf("Current poll standings(%s): \n", poll.ID)

	for _, o := range poll.Options {
		fmt.Printf("%s: %v\n", o.Description, o.Submissions)
	}
}

func SubmitOption(option int, poll *Poll) (PollOption, error) {
	if option >= len(poll.Options) {
		return PollOption{}, fmt.Errorf("wrong submission, length of options too short")
	}

	poll.Options[option].Submissions++
	logPollStandings(poll)

	return poll.Options[option], nil
}
