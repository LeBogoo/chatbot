package chatbot

import (
	"errors"
	"strings"
)

func (chatbot *Chatbot) AddAutoResponse(autoResponse *AutoResponse) error {
	chatbot.AutoResponses = append(chatbot.AutoResponses, autoResponse)
	return nil
}

func (chatbot *Chatbot) GetAutoResponses() []*AutoResponse {
	return chatbot.AutoResponses
}

func (chatbot *Chatbot) RemoveAutoResponse(index int) error {
	if index < 0 || index >= len(chatbot.AutoResponses) {
		return errors.New("index out of range")
	}
	chatbot.AutoResponses = append(chatbot.AutoResponses[:index], chatbot.AutoResponses[index+1:]...)
	return nil
}

func CleanupParts(parts []string) []string {
	cleanedParts := make([]string, 0)
	for _, word := range parts {
		word = strings.ToLower(word)
		word = strings.TrimSpace(word)
		word = strings.Map(func(r rune) rune {
			if r >= 'a' && r <= 'z' {
				return r
			}
			return -1
		}, word)

		cleanedParts = append(cleanedParts, word)
	}

	return cleanedParts
}

func (bot Chatbot) HandleAutoResponse(parts []string, username string) (string, error) {
	cleanedParts := CleanupParts(parts)

	for _, response := range bot.AutoResponses {
		if response.Matches(cleanedParts) {
			return response.Process(cleanedParts, username)
		}
	}

	return "", nil
}
