package gpssh

import (
	"strconv"

	"golang.org/x/crypto/ssh"
)

func Connect(user string, password string, host string, port int) (*ssh.Client, error) {

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host+":"+strconv.Itoa(port), config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Execute(client *ssh.Client, command string) (string, error) {

	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)

	return string(output), err
}
