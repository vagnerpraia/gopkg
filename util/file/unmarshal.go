package gpfile

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Unmarshal(path string, o interface{}) error {

	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		f, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(f, o); err != nil {
			return err
		}
	default:
		return errors.New("error on check extension of file " + path)
	}

	return nil
}
