package chatbot

import (
	"github.com/lebogoo/chatbot/commands"
	"github.com/robfig/cron/v3"
)

type Chatbot struct {
	Prefix               string
	Commands             map[string]commands.Command
	Cron                 *cron.Cron
	Aliases              map[string]string
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
		AutoMessages:        []string{},
		AutoMessageInterval: "*/5 * * * *",
		Cron:                cron.New(),
		autoMessageIndex:    0,
	}
}

func (bot *Chatbot) Start() {
	bot.Cron.AddFunc(bot.AutoMessageInterval, bot.TriggerAutoMessage)
	bot.Cron.Start()
}
