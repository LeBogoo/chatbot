package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lebogoo/chatbot"
	"github.com/lebogoo/chatbot/commands"
	custom "github.com/lebogoo/chatbot/example/commands"
)

func main() {
	bot := chatbot.NewChatbot("!")

	bot.AddCommand(commands.NewSimpleCommand("hello", "Hello, {name}! How are you today, {name}?"))
	bot.AddCommand(commands.NewSimpleCommand("bye", "Goodbye, world!"))
	bot.AddCommand(custom.NewTimeCommand())
	bot.AddAlias("goodbye", "bye")

	bot.AddAutoMessage("Make sure to follow me on GitHub @lebogoo")
	bot.AddAutoMessage("You can also follow me on Instagram @lebogooo")

	bot.AddAutoResponse(chatbot.NewAutoResponse([]string{"what", "your", "instagram"}, "My Instagram handle is @lebogooo"))

	bot.AutoMessageInterval = "* * * * *"
	bot.Start()

	bot.OnAutoMessage(func(message string) {
		fmt.Println("\rBot:", message)
		fmt.Print("You: ")
	})

	commands := bot.ListCommands()
	var commandNames []string
	for name := range commands {
		commandNames = append(commandNames, bot.Prefix+name)
	}

	fmt.Printf("Available commands: %s\n", strings.Join(commandNames, ", "))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')

		response, err := bot.HandleMessage("User", input, false)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if response != "" {
			fmt.Println("Bot:", response)
		}
	}
}
