package gpcommand

type Process struct {
	Executable string
	Envs       []Env
	Args       []Arg
}

func NewProcess(executable string, envs []Env, args []Arg) *Process {

	return &Process{
		Executable: executable,
		Envs:       envs,
		Args:       args,
	}
}
