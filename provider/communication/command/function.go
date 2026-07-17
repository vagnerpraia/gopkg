package gpcommand

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func Run(timeout time.Duration, command string, args ...string) (*Result, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("command %q timed out after %s: %w", command, timeout, ctx.Err())
	}

	result := &Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}

	if err != nil {
		return result, fmt.Errorf("command %q failed: %w", command, err)
	}

	return result, nil
}

func Exists(command string) bool {

	_, err := exec.LookPath(command)

	return err == nil
}
