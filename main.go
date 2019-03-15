package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Text         string       `json:"text"`
	ResponseType string       `json:"response_type"`
	Attachments  []attachment `json:"attachments"`
}

type attachment struct {
	Fallback string   `json:"fallback"`
	Actions  []action `json:"actions"`
}

type action struct {
	Type string `json:"type"`
	Text string `json:"text"`
	URL  string `json:"url"`
}

func HandleRequest() (response, error) {
	p, err := randomProverb()

	resp := response{
		Text:         p.quote,
		ResponseType: "ephemeral",
		Attachments: []attachment{{
			Fallback: p.url,
			Actions: []action{{
				Type: "button",
				Text: "Tell me more... 🎥",
				URL:  p.url,
			}},
		}},
	}

	return resp, err
}

func main() {
	lambda.Start(HandleRequest)
}
