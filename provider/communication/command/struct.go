package gpcommand

import "time"

type Command struct {
	Timeout time.Duration
}

func NewCommand(timeout time.Duration) (*Command, error) {

	return &Command{
		Timeout: timeout,
	}, nil
}

func (that *Command) Run(command string) error {

	return Run(command, that.Timeout)
}
