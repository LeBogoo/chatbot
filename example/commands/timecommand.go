package commands

import (
	"fmt"
	"time"
)

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
