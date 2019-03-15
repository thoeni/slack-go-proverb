package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
{
    "text": "Errors are values.",
    "response_type": "ephemeral",
    "attachments": [
        {
            "fallback": "https://flights.example.com/book/r123456",
            "actions": [
                {
                    "type": "button",
                    "text": "Tell me more... 📹",
                    "url": "https://flights.example.com/book/r123456"
                }
            ]
        }
    ]
}
*/

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

func main() {
	http.HandleFunc("/proverb", proverbHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func proverbHandler(w http.ResponseWriter, r *http.Request) {
	p := randomProverb()

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

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}