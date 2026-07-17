package gpcommand

import (
	"time"
)

type Command struct {
	timeout time.Duration
	print   bool
}

func NewCommand(timeout time.Duration, print bool) (*Command, error) {

	return &Command{
		timeout: timeout,
		print:   print,
	}, nil
}

func (that *Command) Run(command string, args ...string) (*Result, error) {

	return Run(that.timeout, that.print, command, args...)
}

func (that *Command) Exists(command string) bool {

	return Exists(command)
}
