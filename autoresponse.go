package chatbot

import (
	"strings"
)

type AutoResponse struct {
	TriggerWords []string
	Response     string
}

func NewAutoResponse(triggerWords []string, response string) *AutoResponse {
	cleanedTriggerWords := CleanupParts(triggerWords)

	return &AutoResponse{
		TriggerWords: cleanedTriggerWords,
		Response:     response,
	}
}

func (resp *AutoResponse) Matches(parts []string) bool {
	for _, trigger := range resp.TriggerWords {
		found := false
		for _, part := range parts {
			if part == trigger {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func (resp *AutoResponse) Process(parts []string, username string) (string, error) {
	reponse := strings.ReplaceAll(resp.Response, "{username}", username)

	return reponse, nil
}
