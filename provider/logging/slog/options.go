package gpslog

import (
	"fmt"
	"reflect"

	gpconfig "github.com/vagnerpraia/gopkg/util/config"
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
	if err := gpconfig.Unmarshal(filePath, options); err != nil {
		return nil, fmt.Errorf("gpconfig.Unmarshal failed: %w", err)
	}

	options.normalizeOS()

	return &options.Gopkg.Provider.Logging, nil
}

func (that *BaseOptions) normalizeOS() {

	if that.Gopkg.Provider.Logging.OS == "" {
		that.Gopkg.Provider.Logging.OS = "local"
	}
}

func (c *LoggingOptions) IsEmpty() bool {

	return reflect.DeepEqual(LoggingOptions{}, *c)
}
