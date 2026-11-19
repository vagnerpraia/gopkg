package gpcommand

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	gpfilesystem "github.com/vagnerpraia/gopkg/util/filesystem"
)

func (that *Client) ExistsDir(ctx context.Context, dir string) (bool, error) {

	err := that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("test -d %s", dir))
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

func (that *Client) CreateDir(ctx context.Context, dir string) error {

	if err := that.Mkdir(ctx, dir); err != nil {
		return err
	}

	if err := that.SetOwnerAndPermissions(ctx, dir); err != nil {
		return err
	}

	return nil
}

func (that *Client) SetOwnerAndPermissions(ctx context.Context, dir string) error {

	if err := that.Chown(ctx, dir); err != nil {
		return err
	}

	if err := that.Chmod(ctx, dir, "755"); err != nil {
		return err
	}

	return nil
}

func (that *Client) DeployDir(ctx context.Context, dir string) error {

	lastDir := dir

	for _, parentDir := range gpfilesystem.ParentDirs(dir, that.config.OS) {
		exists, err := that.ExistsDir(ctx, parentDir)
		if err != nil {
			return err
		}

		if !exists {
			if err := that.CreateDir(ctx, parentDir); err != nil {
				return err
			}
		}

		lastDir = parentDir
	}

	if err := that.SetOwnerAndPermissions(ctx, lastDir); err != nil {
		return err
	}

	return nil
}

func (that *Client) Mkdir(ctx context.Context, dir string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("sudo mkdir -p %s", dir))
}

func (that *Client) Chown(ctx context.Context, dir string) error {

	command := "sudo chown "

	if that.config.SetRecursiveOwner {
		command += "-R "
	}

	command += "%s:%s %s"

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf(command, that.config.User.Owner, that.config.User.Group, dir))
}

func (that *Client) Chmod(ctx context.Context, dir string, permission string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("sudo chmod %s %s", permission, dir))
}

func (that *Client) Rm(ctx context.Context, dir string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("rm -- %s", dir))
}

func (that *Client) Cleanup(ctx context.Context, dir string) error {

	return that.SSH(ctx, that.config.Host.Name, fmt.Sprintf("rm -rf -- %s/*", dir))
}
