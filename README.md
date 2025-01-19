# Chatbot Library

This is a simple chatbot library implemented in Go. The library supports commands, aliases, and automated messages.

## Features

- Add and remove commands
- Add and remove aliases for commands
- Automated messages at regular intervals
- Support for placeholders in command responses
- Automatic responses for specific keywords

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
    "time"

    "github.com/lebogoo/chatbot"
    "github.com/lebogoo/chatbot/commands"
    custom "github.com/lebogoo/chatbot/example/commands"
)

func main() {
    bot := chatbot.NewChatbot("!")

    bot.AddCommand(commands.NewSimpleCommand("hello", "Hello, {name}! How are you today, {name2}?"))
    bot.AddCommand(commands.NewSimpleCommand("bye", "Goodbye, world!"))
    bot.AddCommand(custom.NewTimeCommand())
    bot.AddAlias("goodbye", "bye")

    bot.AddAutoMessage("Make sure to follow me on GitHub @lebogoo")
    bot.AddAutoMessage("You can also follow me on Instagram @lebogooo")

    bot.AddAutoResponse(chatbot.NewAutoResponse([]string{"what", "your", "instagram"}, "My Instagram handle is @lebogooo"))

    bot.AutoMessageInterval = "*/5 * * * *" // Every 5 minutes
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

func (c *TimeCommand) Execute(args []string, isAdmin bool) (string, error) {
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

To customize the interval at which the automated messages are sent, you can set the `AutoMessageInterval` property to any valid cron expression.

For example:

```go
bot := chatbot.NewChatbot("!")

bot.AddAutoMessage("Remember to stay hydrated!")

bot.AutoMessageInterval = "*/5 * * * *" // Every 5 minutes
bot.Start()
```

### Auto Responses

You can add auto responses that the bot will use to respond to specific messages. For example:

```go
bot.AddAutoResponse(chatbot.NewAutoResponse([]string{"what", "your", "instagram"}, "My Instagram handle is @lebogooo"))
```

To get the list of auto responses, you can use the `GetAutoResponses` method:

```go
autoResponses := bot.GetAutoResponses()
for index, response := range autoResponses {
    fmt.Println(index, response.TriggerWords, response.Response)
}
```

To remove an auto response, you can use the `RemoveAutoResponse` method:

```go
err := bot.RemoveAutoResponse(index)
if err != nil {
    fmt.Println("Error:", err)
}
```

### Placeholders in Command Responses

You can use placeholders in command responses to dynamically insert values. Placeholders are defined using curly braces `{}`.

For example:

```go
bot.AddCommand(commands.NewSimpleCommand("greet", "Hello, {name}!"))
```

When the command is executed, the placeholder `{name}` will be replaced with the corresponding argument provided by the user.

### Handling Messages with User Information

You can handle messages with additional user information such as username and badges. For example:

```go
// A message handler function as might be used in a Twitch chat bot
func OnMessage(username string, badges map[string]bool, message string) {
    isAdmin := badges["broadcaster"] || badges["moderator"]

    response, err := bot.HandleMessage(username, message, isAdmin)
    ...
}
```

## License

This project is licensed under the GNU General Public License (GPL). See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.
