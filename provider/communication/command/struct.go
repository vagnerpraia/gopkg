package gpcommand

import (
	"time"
)

type Command struct {
	timeout time.Duration
}

func NewCommand(timeout time.Duration) (*Command, error) {

	return &Command{
		timeout: timeout,
	}, nil
}

func (that *Command) Run(command string, args ...string) (*Result, error) {

	return Run(that.timeout, command, args...)
}

func (that *Command) Exists(command string) bool {

	return Exists(command)
}
