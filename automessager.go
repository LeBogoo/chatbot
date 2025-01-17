package chatbot

import "time"

func (bot *Chatbot) OnAutoMessage(listener func(string)) {
	bot.AutoMessageListeners = append(bot.AutoMessageListeners, listener)
}

func (bot *Chatbot) autoMessageLoop() {
	for {
		time.Sleep(bot.AutoMessageInterval)

		if len(bot.AutoMessages) == 0 {
			continue
		}

		if bot.autoMessageIndex >= len(bot.AutoMessages) {
			bot.autoMessageIndex = 0
		}

		message := bot.AutoMessages[bot.autoMessageIndex]
		for _, listener := range bot.AutoMessageListeners {
			listener(message)
		}

		bot.autoMessageIndex++
	}
}

func (bot *Chatbot) AddAutoMessage(message string) {
	bot.AutoMessages = append(bot.AutoMessages, message)
}

func (bot *Chatbot) RemoveAutoMessage(index int) {
	if index < 0 || index >= len(bot.AutoMessages) {
		return
	}

	bot.AutoMessages = append(bot.AutoMessages[:index], bot.AutoMessages[index+1:]...)
}
