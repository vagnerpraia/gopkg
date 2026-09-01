package gpcommand

type Env struct {
	Name  string
	Value string
}

func NewEnv(name string, value string) *Env {

	return &Env{
		Name:  name,
		Value: value,
	}
}
