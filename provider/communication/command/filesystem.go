package gpcommand

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	gpfilesystem "github.com/vagnerpraia/gopkg/util/filesystem"
)

func (that *Client) ExistsDir(ctx context.Context, path string) (bool, error) {

	err := that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("test -d %s", path))
	if err == nil {
		return true, nil
	}

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		return false, err
	}

	if exitErr.ExitCode() == 1 {
		return false, nil
	}

	return false, err
}

func (that *Client) CreateDir(ctx context.Context, path string) error {

	if err := that.Mkdir(ctx, path); err != nil {
		return err
	}

	if err := that.SetOwnerAndPermissions(ctx, path); err != nil {
		return err
	}

	return nil
}

func (that *Client) SetOwnerAndPermissions(ctx context.Context, path string) error {

	if err := that.Chown(ctx, path); err != nil {
		return err
	}

	if err := that.Chmod(ctx, path, "755"); err != nil {
		return err
	}

	return nil
}

func (that *Client) DeployDir(ctx context.Context, path string) error {

	for _, p := range gpfilesystem.ParentDirs(path, that.config.OS) {
		exists, err := that.ExistsDir(ctx, p)
		if err != nil {
			return err
		}

		if !exists {
			if err := that.CreateDir(ctx, p); err != nil {
				return err
			}
		} else {
			if err := that.SetOwnerAndPermissions(ctx, path); err != nil {
				return err
			}
		}
	}

	return nil
}

func (that *Client) Mkdir(ctx context.Context, path string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("sudo mkdir -p %s", path))
}

func (that *Client) Chown(ctx context.Context, path string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("sudo chown -R %s:%s %s", that.config.User.Owner, that.config.User.Group, path))
}

func (that *Client) Chmod(ctx context.Context, path string, permission string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("sudo chmod %s %s", permission, path))
}

func (that *Client) Rm(ctx context.Context, path string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("rm -- %s", path))
}

func (that *Client) Cleanup(ctx context.Context, path string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("rm -rf -- %s/*", path))
}
