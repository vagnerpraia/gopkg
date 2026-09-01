package gpcommand

import (
	"time"

	gpenum "github.com/vagnerpraia/gopkg/enum"
)

type Config struct {
	Host    Host
	User    User
	Timeout time.Duration
	OS      gpenum.OS
	Verbose bool
}

type User struct {
	Owner string
	Group string
}

type Host struct {
	Name    string
	Address string
	Port    string
}

func NewConfig(host Host, user User, timeout time.Duration, os gpenum.OS, verbose bool) *Config {

	return &Config{
		Host:    host,
		User:    user,
		Timeout: timeout,
		OS:      os,
		Verbose: verbose,
	}
}

func NewUser(owner string, group string) *User {

	return &User{
		Owner: owner,
		Group: group,
	}
}

func NewHost(name string, address string, port string) *Host {

	return &Host{
		Name:    name,
		Address: address,
		Port:    port,
	}
}
