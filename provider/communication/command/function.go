package gpcommand

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"time"
)

func Run(command string, timeout time.Duration) error {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, command)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("timeout")
	}

	if err != nil {
		return err
	}

	return nil
}
