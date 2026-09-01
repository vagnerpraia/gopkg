package gpcommand

import (
	"errors"
	"strings"
)

type ArgType string

const (
	argTypeCommandAlias = "command"
	argTypeFlagAlias    = "flag"
	argTypeDataAlias    = "data"

	ArgTypeCommand ArgType = argTypeCommandAlias
	ArgTypeFlag    ArgType = argTypeFlagAlias
	ArgTypeData    ArgType = argTypeDataAlias
)

func NewArgType(str string) (ArgType, error) {

	str = strings.ToLower(str)

	switch str {
	case argTypeCommandAlias:
		return ArgTypeCommand, nil

	case argTypeFlagAlias:
		return ArgTypeFlag, nil

	case argTypeDataAlias:
		return ArgTypeData, nil
	}

	return "", errors.New("arg type not found")
}

func (that ArgType) IsCommand() bool {

	return that == ArgTypeCommand
}

func (that ArgType) IsFlag() bool {

	return that == ArgTypeFlag
}

func (that ArgType) IsData() bool {

	return that == ArgTypeData
}

func (that ArgType) String() string {

	switch that {
	case ArgTypeCommand:
		return argTypeCommandAlias

	case ArgTypeFlag:
		return argTypeFlagAlias

	case ArgTypeData:
		return argTypeDataAlias
	}

	return ""
}
