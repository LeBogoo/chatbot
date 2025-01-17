package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lebogoo/chatbot"
	"github.com/lebogoo/chatbot/commands"
	custom "github.com/lebogoo/chatbot/example/commands"
)

func main() {
	bot := chatbot.NewChatbot("!")

	bot.AddCommand(commands.NewSimpleCommand("hello", "Hello, world!"))
	bot.AddCommand(commands.NewSimpleCommand("bye", "Goodbye, world!"))
	bot.AddCommand(custom.NewTimeCommand())
	bot.AddAlias("goodbye", "bye")

	bot.AddAutoMessage("Make sure to follow me on GitHub @lebogoo")
	bot.AddAutoMessage("You can also follow me on Instagram @lebogooo")

	bot.AutoMessageInterval = 10 * time.Second
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
		isCommand := bot.IsCommand(input)
		if !isCommand {
			fmt.Printf("Please enter a valid command with the prefix \"%s\"\n", bot.Prefix)
			continue
		}

		response, err := bot.HandleMessage(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if response != "" {
			fmt.Println("Bot:", response)
		}
	}
}
