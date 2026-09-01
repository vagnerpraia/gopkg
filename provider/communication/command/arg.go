package gpcommand

type Arg struct {
	Name    string
	ArgType ArgType
	Value   string
}

func NewArgCommand(name string) *Arg {

	return &Arg{
		Name:    name,
		ArgType: ArgTypeCommand,
	}
}

func NewArgFlag(name string, value string) *Arg {

	return &Arg{
		Name:    name,
		ArgType: ArgTypeFlag,
		Value:   value,
	}
}

func NewArgData(name string, value string) *Arg {

	return &Arg{
		Name:    name,
		ArgType: ArgTypeData,
		Value:   value,
	}
}

func (that Arg) Same(arg Arg) bool {

	return that.Name == arg.Name &&
		that.ArgType == arg.ArgType
}
