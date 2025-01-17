package commands

type SimpleCommand struct {
	name     string
	response string
}

func NewSimpleCommand(name, response string) *SimpleCommand {
	return &SimpleCommand{name: name, response: response}
}

func (c *SimpleCommand) Execute(args []string) (string, error) {
	return c.response, nil
}

func (c *SimpleCommand) Name() string {
	return c.name
}
