package gpcommand

import (
	"context"
)

func (that *Client) SSH(ctx context.Context, host string, command string) error {

	_, err := that.RunRemote(ctx, "ssh", []string{host, command})
	if err != nil {
		return err
	}

	return nil
}

func (that *Client) SCP(ctx context.Context, args ...string) error {

	_, err := that.RunRemote(ctx, "scp", args)
	if err != nil {
		return err
	}

	return nil
}

func (that *Client) Git(ctx context.Context, args ...string) (*Result, error) {

	result, err := that.RunRemote(ctx, "git", args)
	if err != nil {
		return nil, err
	}

	return result, nil
}
