package gpcommand

import (
	"context"
	"strings"

	gpfilesystem "github.com/vagnerpraia/gopkg/util/filesystem"
)

func (that *Client) ExecRemote(ctx context.Context, pathBin string, process Process) error {

	parts := []string{}
	for _, env := range process.Envs {
		parts = append(parts, env.Name+"='"+env.Value+"'")
	}

	parts = append(parts, "exec")

	pathBin = gpfilesystem.NormalizePathOS(pathBin+"/"+process.Executable, that.config.OS)

	parts = append(parts, pathBin)

	parts = append(parts, "run")

	for _, arg := range process.Args {
		switch arg.ArgType {
		case ArgTypeCommand:
			parts = append(parts, arg.Name)

		case ArgTypeFlag:
			if arg.Value == "" {
				parts = append(parts, "-"+arg.Name)
			} else {
				parts = append(parts, "-"+arg.Name+"='"+arg.Value+"'")
			}

		case ArgTypeData:
			if arg.Value != "" {
				parts = append(parts, "'"+arg.Value+"'")
			}
		}
	}

	command := strings.Join(parts, " ")

	return that.SSH(ctx, that.config.Host.Name, command)
}

func (that *Client) ExecLocal(ctx context.Context, process Process) (*Result, error) {

	envs := []string{}
	args := []string{}

	for _, env := range process.Envs {
		envs = append(envs, env.Name+"="+env.Value)
	}

	for _, arg := range process.Args {
		switch arg.ArgType {
		case ArgTypeCommand:
			args = append(args, arg.Name)

		case ArgTypeFlag:
			if arg.Value == "" {
				args = append(args, "-"+arg.Name)
			} else {
				args = append(args, "-"+arg.Name+"="+arg.Value)
			}

		case ArgTypeData:
			if arg.Value != "" {
				args = append(args, arg.Value)
			}
		}
	}

	result, err := that.RunLocal(ctx, process.Executable, envs, args)
	if err != nil {
		return nil, err
	}

	return result, nil
}
