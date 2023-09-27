package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Job func()

type JobCreator interface {
	GoogleChatJob(webhook, message string) Job
}

func NewJobCreator() JobCreator {
	return &jobCreator{}
}

type jobCreator struct{}

func (jobCreator) GoogleChatJob(webhook, message string) Job {
	return func() {
		log.Printf(`{"msg":"send google chat","body":"%s"}`, message)
		payload, err := json.Marshal(struct {
			Text string `json:"text"`
		}{
			Text: message,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		resp, err := http.Post(webhook, "application/json; charset=UTF-8", bytes.NewReader(payload))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("HTTP: %v\n", resp.StatusCode)
		}
	}
}

func NewJobCreatorMock() JobCreator {
	return &jobCreatorMock{}
}

type jobCreatorMock struct{}

func (jobCreatorMock) GoogleChatJob(webhook, message string) Job {
	return func() {
		fmt.Println(message)
	}
}
