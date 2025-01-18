# Chatbot Library

This is a simple chatbot library implemented in Go. The library supports commands, aliases, and automated messages.

## Features

- Add and remove commands
- Add and remove aliases for commands
- Automated messages at regular intervals
- Support for placeholders in command responses

## Installation

To use this library in your project, you can get it using:

```sh
go get github.com/lebogoo/chatbot
```

Then, import it in your Go code:

```go
import "github.com/lebogoo/chatbot"
```

## Usage

Here is an example of how to use the chatbot library:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/lebogoo/chatbot"
    "github.com/lebogoo/chatbot/commands"
)

func main() {
    bot := chatbot.NewChatbot("!")

    bot.AddCommand(commands.NewSimpleCommand("hello", "Hello, {name}! How are you today, {name2}?"))
    bot.AddCommand(commands.NewSimpleCommand("bye", "Goodbye, world!"))
    bot.AddAlias("goodbye", "bye")

    bot.AddAutoMessage("Make sure to follow me on GitHub @lebogoo")
    bot.AddAutoMessage("You can also follow me on Instagram @lebogooo")

    bot.AutoMessageInterval = 30 * time.Second
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
```

### Adding Commands

You can add new commands by using the `AddCommand` method.
There is a `SimpleCommand` struct that you can use to create a new "simple" command.

For example:

```go
bot.AddCommand(commands.NewSimpleCommand("newcommand", "This is a new command!"))
```

In order to create custom commands, you can implement the `Command` interface.
The `Command` interface has two methods: `Execute` and `Name`.

Here is an example of a custom command:

```go
type TimeCommand struct {
}

func NewTimeCommand() *TimeCommand {
	return &TimeCommand{}
}

func (c *TimeCommand) Execute(args []string) (string, error) {
	return fmt.Sprintf("The current time is: %s", time.Now().Format("15:04:05")), nil
}

func (c *TimeCommand) Name() string {
	return "time"
}
```

### Adding Aliases

You can add aliases for existing commands by using the `AddAlias` method. For example:

```go
bot.AddAlias("aliasname", "existingcommand")
```

### Automated Messages

You can add automated messages that the bot will send at regular intervals. For example:

```go
bot.AddAutoMessage("Remember to stay hydrated!")
```

To customize the interval at which the automated messages are sent, you can set the `AutoMessageInterval` property.

For example:

```go
bot := chatbot.NewChatbot("!")

bot.AddAutoMessage("Remember to stay hydrated!")

bot.AutoMessageInterval = 10 * time.Second
bot.Start()
```

### Placeholders in Command Responses

You can use placeholders in command responses to dynamically insert values. Placeholders are defined using curly braces `{}`.

For example:

```go
bot.AddCommand(commands.NewSimpleCommand("greet", "Hello, {name}!"))
```

When the command is executed, the placeholder `{name}` will be replaced with the corresponding argument provided by the user.

## License

This project is licensed under the GNU General Public License (GPL). See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.
