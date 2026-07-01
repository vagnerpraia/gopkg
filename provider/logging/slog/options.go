package gpslog

import (
	"fmt"
	"reflect"

	gpfile "github.com/vagnerpraia/gopkg/util/file"
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
	PathFile string `yaml:"pathFile"`
}

func NewLoggingOptions(filePath string) (*LoggingOptions, error) {

	options := &BaseOptions{}
	if err := gpfile.Unmarshal(filePath, options); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &options.Gopkg.Provider.Logging, nil
}

func (c *LoggingOptions) IsEmpty() bool {

	return reflect.DeepEqual(LoggingOptions{}, *c)
}
