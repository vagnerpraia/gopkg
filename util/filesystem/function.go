package gpfilesystem

import (
	"path"
	"path/filepath"
	"strings"

	gpenum "github.com/vagnerpraia/gopkg/enum"
)

func NormalizePath(str string) string {

	return path.Clean(filepath.ToSlash(str))
}

func NormalizePathLocal(str string) string {

	return NormalizePath(str)
}

func NormalizePathOS(str string, os gpenum.OS) string {

	str = NormalizePath(str)

	if os.IsLocal() {
		return str
	}

	if os.IsWindows() {
		str = strings.ReplaceAll(str, "/", `\`)
	}

	return str
}
