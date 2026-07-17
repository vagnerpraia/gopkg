package gpcommand

import (
	"time"
)

type Command struct {
	Timeout time.Duration
}

func NewCommand(timeout time.Duration) (*Command, error) {

	return &Command{
		Timeout: timeout,
	}, nil
}

func (that *Command) Run(command string, args ...string) (*Result, error) {

	return Run(that.Timeout, command, args...)
}

func (that *Command) Exists(command string) bool {

	return Exists(command)
}
