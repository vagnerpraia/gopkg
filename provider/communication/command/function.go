package gpcommand

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(ctx context.Context, timeout time.Duration, command string, args ...string) (*Result, error) {

	commandCtx := ctx

	if timeout > 0 {
		var cancel context.CancelFunc
		commandCtx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(commandCtx, command, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := &Result{
		Command: strings.Join(append([]string{command}, args...), " "),
		Stdout:  stdout.String(),
		Stderr:  stderr.String(),
	}

	switch {
	case errors.Is(commandCtx.Err(), context.DeadlineExceeded):
		return result, fmt.Errorf("command %q timed out after %s: %w", result.Command, timeout, commandCtx.Err())

	case errors.Is(commandCtx.Err(), context.Canceled):
		return result, fmt.Errorf("command %q was canceled: %w", result.Command, commandCtx.Err())
	}

	if err != nil {
		if result.Stderr != "" {
			return result, fmt.Errorf("command %q failed: %w\nstderr:\n%s", result.Command, err, result.Stderr)
		}

		return result, fmt.Errorf("command %q failed: %w", result.Command, err)
	}

	return result, nil
}

func Exists(command string) bool {

	_, err := exec.LookPath(command)

	return err == nil
}
