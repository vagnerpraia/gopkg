package gpcommand

import (
	"context"
)

func (that *Client) GitCommit(ctx context.Context, repository string) (string, error) {

	result, err := that.Git(ctx, "-C", repository, "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}

	return result.Stdout, nil
}
