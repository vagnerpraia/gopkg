package gpmongodb

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
	MongoDB MongoDBOptions `yaml:"mongodb"`
}

type MongoDBOptions struct {
	Client ClientOptions `yaml:"client"`
}

type ClientOptions struct {
	URI string `yaml:"uri"`
	DB  string `yaml:"db"`
}

func NewClientOptions(filePath string) (*ClientOptions, error) {

	options := &BaseOptions{}
	if err := gpfile.Unmarshal(filePath, options); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &options.Gopkg.Provider.MongoDB.Client, nil
}

func (c *ClientOptions) IsEmpty() bool {

	return reflect.DeepEqual(ClientOptions{}, *c)
}
