package gpfilesystem

import (
	"path/filepath"
)

func NormalizePath(path string) string {

	return filepath.Clean(filepath.ToSlash(path))
}
