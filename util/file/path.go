package gpfile

import (
	"runtime"
	"strings"
)

func NormalizePath(path string) string {

	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "/", "\\")
	} else {
		path = strings.ReplaceAll(path, "\\", "/")
	}

	return path
}
