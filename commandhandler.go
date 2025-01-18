package chatbot

import (
	"regexp"
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

func (bot Chatbot) HandleCommand(name string, args []string, username string, isAdmin bool) (string, error) {
	name = cleanupName(name)

	if alias, exists := bot.Aliases[name]; exists {
		name = alias
	}

	if command, exists := bot.Commands[name]; exists {
		response, err := command.Execute(username, args, isAdmin)
		if err != nil {
			return "", err
		}

		processedResponse, err := bot.postProcessResponse(response, name, args, username)
		if err != nil {
			return bot.CreateUsage(response, name), nil
		}

		return processedResponse, nil
	}

	return "", commands.ErrCommandNotFound
}

func cleanupName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

func (chatbot Chatbot) IsCommand(message string) bool {
	return message[:len(chatbot.Prefix)] == chatbot.Prefix
}

func (chatbot Chatbot) HandleMessage(username, message string, isAdmin bool) (string, error) {
	if message[:len(chatbot.Prefix)] != chatbot.Prefix {
		return "", commands.ErrInvalidPrefix
	}

	parts := strings.Split(message, " ")

	return chatbot.HandleCommand(parts[0][len(chatbot.Prefix):], parts[1:], username, isAdmin)
}

func (chatbot Chatbot) CreateUsage(response string, command string) string {
	response = strings.ReplaceAll(response, "{command}", command)

	re := regexp.MustCompile(`\{(\w+)\}`)
	matches := re.FindAllStringSubmatch(response, -1)

	placeholderIndex := make(map[string]int)
	for i, match := range matches {
		placeholderIndex[match[1]] = i
	}

	usage := "Usage: " + chatbot.Prefix + command
	for name := range placeholderIndex {
		usage += " <" + name + ">"
	}

	return usage
}

func (chatbot Chatbot) postProcessResponse(response string, command string, args []string, username string) (string, error) {
	response = strings.ReplaceAll(response, "{command}", command)
	response = strings.ReplaceAll(response, "{username}", username)

	// get all named placeholders like {name} using regex
	re := regexp.MustCompile(`\{(\w+)\}`)
	matches := re.FindAllStringSubmatch(response, -1)

	uniqueMatches := make(map[string]string)
	for _, match := range matches {
		uniqueMatches[match[1]] = match[1]
	}

	if len(uniqueMatches) > len(args) {
		return "", commands.ErrMissingParameters
	}

	placeholderMap := make(map[string]string)

	// get keys of the unique matches
	var keys []string
	for key := range uniqueMatches {
		keys = append(keys, key)
	}

	// create a map of the unique match and their values
	// explicitly use the unique map to avoid duplicate keys
	for index := range len(uniqueMatches) {
		placeholderMap[keys[index]] = strings.TrimSpace(args[index])
	}

	// replace all placeholders with their values
	for placeholder, value := range placeholderMap {
		response = strings.ReplaceAll(response, "{"+placeholder+"}", value)
	}

	return response, nil
}
