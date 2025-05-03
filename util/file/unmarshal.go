package gpfile

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/mitchellh/mapstructure"
)

func Unmarshal(path string, output interface{}) error {

	paths := []string{path}
	if strings.Contains(path, ",") {
		paths = strings.Split(path, ",")
	} else if strings.Contains(path, ";") {
		paths = strings.Split(path, ";")
	}

	merged := make(map[string]interface{})

	for _, path := range paths {
		fileContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var current map[string]interface{}
		if err := yaml.Unmarshal(fileContent, &current); err != nil {
			return err
		}

		mergeMaps(merged, current)
	}

	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			Result:  output,
			TagName: "yaml",
		},
	)
	if err != nil {
		return err
	}

	if err := decoder.Decode(merged); err != nil {
		return err
	}

	return nil
}

func mergeMaps(dest map[string]interface{}, src map[string]interface{}) {

	for key, srcVal := range src {
		if destVal, ok := dest[key]; ok {
			destMap, destIsMap := destVal.(map[string]interface{})
			srcMap, srcIsMap := srcVal.(map[string]interface{})

			if destIsMap && srcIsMap {
				mergeMaps(destMap, srcMap)
				continue
			}
		}

		dest[key] = srcVal
	}
}
