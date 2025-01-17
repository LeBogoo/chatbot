package chatbot

import (
	"time"

	"github.com/lebogoo/chatbot/commands"
)

type Chatbot struct {
	Prefix               string
	Commands             map[string]commands.Command
	Aliases              map[string]string
	AutoMessages         []string
	AutoMessageListeners []func(string)
	AutoMessageInterval  time.Duration
	autoMessageIndex     int
}

func NewChatbot(prefix string) Chatbot {
	return Chatbot{
		Prefix:              prefix,
		Commands:            make(map[string]commands.Command),
		Aliases:             make(map[string]string),
		AutoMessages:        []string{},
		AutoMessageInterval: 60 * time.Second,
		autoMessageIndex:    0,
	}
}

func (bot *Chatbot) Start() {
	go bot.autoMessageLoop()
}
