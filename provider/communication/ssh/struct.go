package gpssh

import (
	"golang.org/x/crypto/ssh"
)

type SSH struct {
	Config Config
	Client *ssh.Client
}

func NewSSH(host string, port int, user string, password string) (*SSH, error) {

	auth := AuthConfig{
		User:     user,
		Password: password,
	}

	config := Config{
		Host: host,
		Port: port,
		Auth: auth,
	}

	return &SSH{
		Config: config,
	}, nil
}

func (that *SSH) Connect(cmd string) (*ssh.Client, error) {

	return Connect(
		that.Config.Auth.User,
		that.Config.Auth.Password,
		that.Config.Host,
		that.Config.Port,
	)
}

func (that *SSH) Execute(cmd string) (string, error) {

	return Execute(that.Client, cmd)
}

func (that *SSH) ConnectExecute(cmd string) (string, error) {

	client, err := Connect(
		that.Config.Auth.User,
		that.Config.Auth.Password,
		that.Config.Host,
		that.Config.Port,
	)

	if err != nil {
		return "", err
	}

	that.Client = client

	return that.Execute(cmd)
}
