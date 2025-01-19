package chatbot

import (
	"strings"

	"github.com/lebogoo/chatbot/commands"
	"github.com/robfig/cron/v3"
)

type Chatbot struct {
	Prefix               string
	Commands             map[string]commands.Command
	Cron                 *cron.Cron
	Aliases              map[string]string
	AutoResponses        []*AutoResponse
	AutoMessages         []string
	AutoMessageListeners []func(string)
	AutoMessageInterval  string
	autoMessageIndex     int
}

func NewChatbot(prefix string) *Chatbot {
	return &Chatbot{
		Prefix:              prefix,
		Commands:            make(map[string]commands.Command),
		Aliases:             make(map[string]string),
		AutoResponses:       []*AutoResponse{},
		AutoMessages:        []string{},
		AutoMessageInterval: "*/5 * * * *",
		Cron:                cron.New(),
		autoMessageIndex:    0,
	}
}

func (chatbot Chatbot) HandleMessage(username, message string, isAdmin bool) (string, error) {
	parts := strings.Split(message, " ")
	if chatbot.IsCommand(message) {
		return chatbot.HandleCommand(parts[0][len(chatbot.Prefix):], parts[1:], username, isAdmin)
	}

	return chatbot.HandleAutoResponse(parts, username)
}

func (bot *Chatbot) Start() {
	bot.Cron.AddFunc(bot.AutoMessageInterval, bot.TriggerAutoMessage)
	bot.Cron.Start()
}
