package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
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

func HandleRequest() (events.APIGatewayProxyResponse, error) {
	p, err := randomProverb()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    map[string]string{"Content-Type": "application/json; charset=UTF-8"},
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	resp := response{
		Text:         p.quote,
		ResponseType: "in_channel",
		Attachments: []attachment{{
			Fallback: p.url,
			Actions: []action{{
				Type: "button",
				Text: "Tell me more... 🎥",
				URL:  p.url,
			}},
		}},
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    map[string]string{"Content-Type": "application/json; charset=UTF-8"},
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json; charset=UTF-8"},
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
