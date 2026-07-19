package gpcommand

import (
	"context"
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

func (that *Command) Run(ctx context.Context, command string, args ...string) (*Result, error) {

	return Run(ctx, that.timeout, command, args...)
}

func (that *Command) Exists(command string) bool {

	return Exists(command)
}
