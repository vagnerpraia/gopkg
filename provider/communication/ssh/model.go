package gpssh

import (
	"time"
)

type Config struct {
	Host       string
	Port       int
	Timeout    time.Duration
	Auth       AuthConfig
	KnownHosts KnownHostsConfig
}

type AuthConfig struct {
	User       string
	Password   string
	PrivateKey PrivateKeyConfig
}

type PrivateKeyConfig struct {
	File       string
	Passphrase string
}

type KnownHostsConfig struct {
	File string
}
