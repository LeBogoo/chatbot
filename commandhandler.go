package chatbot

import (
	"strings"

	"github.com/lebogoo/chatbot/commands"
)

func (chatbot Chatbot) AddCommand(command commands.Command) error {
	name := cleanupName(command.Name())
	if _, exists := chatbot.Commands[name]; exists {
		return commands.ErrCommandAlreadyExists
	}

	chatbot.Commands[name] = command
	return nil
}

func (chatbot Chatbot) GetCommand(name string) (commands.Command, error) {
	name = cleanupName(name)
	if command, exists := chatbot.Commands[name]; exists {
		return command, nil
	}

	return nil, commands.ErrCommandNotFound
}

func (chatbot Chatbot) UpdateCommand(command commands.Command) error {
	name := cleanupName(command.Name())
	if _, exists := chatbot.Commands[name]; !exists {
		return commands.ErrCommandNotFound
	}

	chatbot.Commands[name] = command
	return nil
}

func (chatbot Chatbot) RemoveCommand(name string) error {
	name = cleanupName(name)
	if _, exists := chatbot.Commands[name]; !exists {
		return commands.ErrCommandNotFound
	}

	// delete all alaises for the command
	for alias, commandName := range chatbot.Aliases {
		if commandName == name {
			delete(chatbot.Aliases, alias)
		}
	}

	delete(chatbot.Commands, name)
	return nil
}

func (chatbot Chatbot) ListCommands() map[string]commands.Command {
	return chatbot.Commands
}

func (chatbot Chatbot) AddAlias(alias, commandName string) error {
	alias = cleanupName(alias)
	commandName = cleanupName(commandName)

	if _, exists := chatbot.Commands[commandName]; !exists {
		return commands.ErrCommandNotFound
	}

	if _, exists := chatbot.Commands[alias]; exists {
		return commands.ErrCommandAlreadyExists
	}

	if _, exists := chatbot.Aliases[alias]; exists {
		return commands.ErrAliasAlreadyExists
	}

	chatbot.Aliases[alias] = commandName
	return nil
}

func (chatbot Chatbot) RemoveAlias(alias string) error {
	alias = cleanupName(alias)
	if _, exists := chatbot.Aliases[alias]; !exists {
		return commands.ErrAliasNotFound
	}

	delete(chatbot.Aliases, alias)
	return nil
}

func (bot Chatbot) HandleCommand(name string, args []string) (string, error) {
	name = cleanupName(name)

	if alias, exists := bot.Aliases[name]; exists {
		name = alias
	}

	if command, exists := bot.Commands[name]; exists {
		return command.Execute(args)
	}

	return "", commands.ErrCommandNotFound
}

func cleanupName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

func (chatbot Chatbot) IsCommand(message string) bool {
	return message[:len(chatbot.Prefix)] == chatbot.Prefix
}

func (chatbot Chatbot) HandleMessage(message string) (string, error) {
	if message[:len(chatbot.Prefix)] != chatbot.Prefix {
		return "", commands.ErrInvalidPrefix
	}

	parts := strings.Split(message, " ")

	return chatbot.HandleCommand(parts[0][len(chatbot.Prefix):], parts[1:])
}
