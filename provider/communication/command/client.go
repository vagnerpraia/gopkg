package gpcommand

import (
	"context"
)

type Client struct {
	envs   []string
	config *Config
}

func NewClient(config *Config) (*Client, error) {

	return &Client{
		config: config,
	}, nil
}

func (that *Client) Verbose() bool {

	return that.config.Verbose
}

func (that *Client) SetEnvs(envs []string) {

	that.envs = envs
}

func (that *Client) RunRemote(ctx context.Context, command string, args []string) (*Result, error) {

	result, err := that.Run(ctx, command, args)

	if that.config.Verbose {
		result.Print()
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (that *Client) RunLocal(ctx context.Context, command string, envs []string, args []string) (*Result, error) {

	that.SetEnvs(envs)

	result, err := that.Run(ctx, command, args)

	if that.config.Verbose {
		result.Print()
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (that *Client) Interactive(
	ctx context.Context, command string, args ...string,
) error {

	// session, err := that.client.NewSession()
	// if err != nil {
	// 	return err
	// }
	// defer session.Close()

	// modes := ssh.TerminalModes{
	// 	ssh.ECHO:          0,
	// 	ssh.TTY_OP_ISPEED: 14400,
	// 	ssh.TTY_OP_OSPEED: 14400,
	// }

	// if err := session.RequestPty("xterm", 24, 80, modes); err != nil {
	// 	return err
	// }

	// session.Stdin = os.Stdin
	// session.Stdout = os.Stdout
	// session.Stderr = os.Stderr

	// cmd := command
	// if len(args) > 0 {
	// 	cmd += " " + strings.Join(args, " ")
	// }

	// if err := session.Start(cmd); err != nil {
	// 	return err
	// }

	// if err := session.Wait(); err != nil {
	// 	return err
	// }

	return nil
}
