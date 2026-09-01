package gpcommand

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
)

func (that *Client) Run(ctx context.Context, command string, args []string) (*Result, error) {

	commandCtx := ctx

	if that.config.Timeout > 0 {
		var cancel context.CancelFunc
		commandCtx, cancel = context.WithTimeout(ctx, that.config.Timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(commandCtx, command, args...)

	cmd.Env = append(os.Environ(), that.envs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := &Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}

	switch {
	case errors.Is(commandCtx.Err(), context.DeadlineExceeded):
		return result, err

	case errors.Is(commandCtx.Err(), context.Canceled):
		return result, err
	}

	if err != nil {
		if result.Stderr != "" {
			return result, err
		}

		return result, err
	}

	return result, nil
}

func (that *Client) Exists(command string) bool {

	_, err := exec.LookPath(command)

	return err == nil
}
