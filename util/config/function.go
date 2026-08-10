package gpconfig

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func Unmarshal(path string, output any) error {

	paths := []string{path}

	if strings.Contains(path, ",") {
		paths = strings.Split(path, ",")
	} else if strings.Contains(path, ";") {
		paths = strings.Split(path, ";")
	}

	var merged yaml.Node

	for _, path := range paths {
		fileContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var current yaml.Node

		if err := yaml.Unmarshal(fileContent, &current); err != nil {
			return err
		}

		if err := mergeNodes(&merged, &current); err != nil {
			return err
		}
	}

	if err := merged.Decode(output); err != nil {
		return err
	}

	return nil
}

func mergeNodes(dest, src *yaml.Node) error {

	if src.Kind == 0 {
		return nil
	}

	if dest.Kind == 0 {
		*dest = *src
		return nil
	}

	destMapping := mappingNode(dest)
	srcMapping := mappingNode(src)

	if destMapping != nil && srcMapping != nil {
		mergeMappingNodes(destMapping, srcMapping)
		return nil
	}

	*dest = *src

	return nil
}

func mappingNode(node *yaml.Node) *yaml.Node {

	if node.Kind == yaml.MappingNode {
		return node
	}

	if node.Kind == yaml.DocumentNode && len(node.Content) == 1 {
		if node.Content[0].Kind == yaml.MappingNode {
			return node.Content[0]
		}
	}

	return nil
}

func mergeMappingNodes(dest, src *yaml.Node) {

	for i := 0; i < len(src.Content); i += 2 {
		srcKey := src.Content[i]
		srcValue := src.Content[i+1]

		destValue := findMappingValue(dest, srcKey.Value)

		if destValue == nil {
			dest.Content = append(dest.Content, srcKey, srcValue)
			continue
		}

		destMap := mappingNode(destValue)
		srcMap := mappingNode(srcValue)

		if destMap != nil && srcMap != nil {
			mergeMappingNodes(destMap, srcMap)
			continue
		}

		*destValue = *srcValue
	}
}

func findMappingValue(mapping *yaml.Node, key string) *yaml.Node {

	if mapping.Kind != yaml.MappingNode {
		return nil
	}

	for i := 0; i < len(mapping.Content); i += 2 {
		if mapping.Content[i].Value == key {
			return mapping.Content[i+1]
		}
	}

	return nil
}
