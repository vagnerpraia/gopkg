package gpfilesystem

import (
	"fmt"
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

func Subdirectories(finalPath string, relativePath string, os gpenum.OS) ([]string, error) {

	finalPath = NormalizePathOS(finalPath, os)
	relativePath = NormalizePathOS(relativePath, os)

	if finalPath == relativePath {
		return []string{finalPath}, nil
	}

	rel, err := filepath.Rel(relativePath, finalPath)
	if err != nil {
		return nil, err
	}

	if rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return nil, fmt.Errorf("%q is not below %q", finalPath, relativePath)
	}

	parts := strings.Split(rel, string(filepath.Separator))

	paths := make([]string, 0, len(parts))
	current := relativePath

	for _, part := range parts {
		current = filepath.Join(current, part)
		paths = append(paths, current)
	}

	return paths, nil
}
