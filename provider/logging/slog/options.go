package gpslog

import (
	"fmt"
	"reflect"

	gpfilesystem "github.com/vagnerpraia/gopkg/util/filesystem"
)

type BaseOptions struct {
	Gopkg GopkgOptions `yaml:"gopkg"`
}

type GopkgOptions struct {
	Provider ProviderOptions `yaml:"provider"`
}

type ProviderOptions struct {
	Logging LoggingOptions `yaml:"logging"`
}

type LoggingOptions struct {
	Print     bool   `yaml:"print"`
	WriteFile bool   `yaml:"writeFile"`
	PathFile  string `yaml:"pathFile"`
	OS        string `yaml:"os"`
}

func NewLoggingOptions(filePath string) (*LoggingOptions, error) {

	options := &BaseOptions{}
	if err := gpfilesystem.Unmarshal(filePath, options); err != nil {
		return nil, fmt.Errorf("gpfilesystem.Unmarshal failed: %w", err)
	}

	return &options.Gopkg.Provider.Logging, nil
}

func (c *LoggingOptions) IsEmpty() bool {

	return reflect.DeepEqual(LoggingOptions{}, *c)
}
