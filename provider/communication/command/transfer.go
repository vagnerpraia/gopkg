package gpcommand

import (
	"context"
	"fmt"
)

func (that *Client) Upload(ctx context.Context, fileLocal string, fileRemote string) error {

	return that.SCP(ctx, fileLocal, fmt.Sprintf("%s:%s", that.config.Host.Name, fileRemote))
}

func (that *Client) Download(ctx context.Context, fileLocal string, fileRemote string) error {

	return that.SCP(ctx, fmt.Sprintf("%s:%s", that.config.Host.Name, fileRemote), fileLocal)
}
