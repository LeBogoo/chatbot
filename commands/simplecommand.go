package commands

type SimpleCommand struct {
	name     string
	response string
}

func NewSimpleCommand(name, response string) *SimpleCommand {
	return &SimpleCommand{name: name, response: response}
}

func (c *SimpleCommand) Execute(username string, args []string, isAdmin bool) (string, error) {
	return c.response, nil
}

func (c *SimpleCommand) Name() string {
	return c.name
}

func (c *SimpleCommand) Response() string {
	return c.response
}
