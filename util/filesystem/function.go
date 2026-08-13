package gpfilesystem

import (
	"fmt"
	"path"
	"path/filepath"
	"slices"
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
		paths = append(paths, NormalizePathOS(current, os))
	}

	return paths, nil
}

func ParentDirs(path string, os gpenum.OS) []string {

	path = NormalizePathOS(path, os)

	dirs := []string{}

	for dir := filepath.Dir(path); dir != path; {
		if dir == filepath.Dir(dir) {
			break
		}

		dirs = append(dirs, NormalizePathOS(dir, os))

		path = dir
		dir = filepath.Dir(path)
	}

	slices.Reverse(dirs)

	return dirs
}

func IsWithin(path string, root string, os gpenum.OS) bool {

	path = NormalizePathOS(path, os)
	root = NormalizePathOS(root, os)

	if path == root {
		return false
	}

	if os.IsWindows() {
		root += "\\"
	} else {
		root += "/"
	}

	return strings.HasPrefix(path, root)
}
