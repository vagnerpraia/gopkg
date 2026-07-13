package gpmongodb

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
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
	Host                   string `yaml:"host"`
	Port                   int    `yaml:"port"`
	DB                     string `yaml:"db"`
	User                   string `yaml:"user"`
	AuthDB                 string `yaml:"authDB"`
	ConnectTimeout         int    `yaml:"connectTimeout"`
	MaxConnIdleTime        int    `yaml:"maxConnIdleTime"`
	MaxPoolSize            uint64 `yaml:"maxPoolSize"`
	MinPoolSize            uint64 `yaml:"minPoolSize"`
	ServerSelectionTimeout int    `yaml:"serverSelectionTimeout"`
}

func NewClientOptions(filePath string) (*ClientOptions, error) {

	options := &BaseOptions{}
	if err := gpfile.Unmarshal(filePath, options); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &options.Gopkg.Provider.MongoDB.Client, nil
}

func (that *ClientOptions) MongoURI() (string, error) {

	u := &url.URL{
		Scheme: "mongodb",
		Host:   net.JoinHostPort(that.Host, fmt.Sprintf("%d", that.Port)),
	}

	if that.DB != "" {
		u.Path = "/" + that.DB
	}

	if that.User != "" {
		password, err := that.Password()
		if err != nil {
			return "", fmt.Errorf("os.Getenv failed: %w", err)
		}

		u.User = url.UserPassword(that.User, password)
	}

	if that.AuthDB != "" {
		query := u.Query()
		query.Set("authSource", that.AuthDB)
		u.RawQuery = query.Encode()
	}

	return u.String(), nil
}

func (that *ClientOptions) Validate() error {

	var errs []error

	if that.Host == "" {
		errs = append(errs, fmt.Errorf("mongodb.host cannot be empty"))
	}

	if that.Port < 1 || that.Port > 65535 {
		errs = append(errs,
			fmt.Errorf("mongodb.port must be between 1 and 65535 (got %d)", that.Port))
	}

	if that.ConnectTimeout <= 0 {
		errs = append(errs, fmt.Errorf("mongodb.connectTimeout must be greater than zero (got %d)", that.ConnectTimeout))
	}

	if that.MaxConnIdleTime <= 0 {
		errs = append(errs, fmt.Errorf("mongodb.maxConnIdleTime must be greater than zero (got %d)", that.MaxConnIdleTime))
	}

	if that.ServerSelectionTimeout <= 0 {
		errs = append(errs, fmt.Errorf("mongodb.serverSelectionTimeout must be greater than zero (got %d)", that.ServerSelectionTimeout))
	}

	if that.MaxPoolSize == 0 {
		errs = append(errs, fmt.Errorf("mongodb.maxPoolSize must be greater than zero"))
	}

	if that.MinPoolSize > that.MaxPoolSize {
		errs = append(errs,
			fmt.Errorf(
				"mongodb.minPoolSize (%d) cannot be greater than mongodb.maxPoolSize (%d)",
				that.MinPoolSize,
				that.MaxPoolSize,
			))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (that *ClientOptions) Password() (string, error) {

	passwordEnv := "INSAIO_QUANTOS_MONGODB_PASSWORD"

	password := os.Getenv(passwordEnv)
	if password == "" {
		return "", fmt.Errorf("environment variable %q not defined", passwordEnv)
	}

	return password, nil
}

func (that *ClientOptions) IsEmpty() bool {

	return reflect.DeepEqual(ClientOptions{}, *that)
}
