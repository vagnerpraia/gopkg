package gpfilesystem

import (
	"errors"
	"strings"
)

type OS string

const (
	osLocalAlias   = "local"
	osLinuxAlias   = "linux"
	osWindowsAlias = "windows"

	OSLocal   OS = osLocalAlias
	OSLinux   OS = osLinuxAlias
	OSWindows OS = osWindowsAlias
)

func NewOS(str string) (OS, error) {

	str = strings.ToLower(str)

	switch str {
	case osLocalAlias:
		return OSLocal, nil

	case osLinuxAlias:
		return OSLinux, nil

	case osWindowsAlias:
		return OSWindows, nil
	}

	return "", errors.New("os not found")
}

func (that OS) IsLocal() bool {

	return that == OSLocal
}

func (that OS) IsLinux() bool {

	return that == OSLinux
}

func (that OS) IsWindows() bool {

	return that == OSWindows
}
