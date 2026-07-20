package gpstring

import (
	"path"
	"path/filepath"
)

func NormalizePath(str string) string {

	return path.Clean(filepath.ToSlash(str))
}
