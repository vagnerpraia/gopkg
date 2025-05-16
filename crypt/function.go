package gpcrypt

import (
	"encoding/base64"
	"hash"
)

func FromInterface(hasher hash.Hash, i interface{}) (string, error) {

	code, err := digest(i, hasher)
	if err != nil {
		return "", err
	}

	return code, nil
}

func FromInterfaces(hasher hash.Hash, i ...interface{}) (string, error) {

	code := ""
	for _, ii := range i {
		s, err := digest(ii, hasher)
		if err != nil {
			return "", err
		}

		code += s
	}

	hasher.Write([]byte(code))

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil)), nil
}

func FromString(hasher hash.Hash, str string) (string, error) {

	code, err := digest(str, hasher)
	if err != nil {
		return "", err
	}

	return code, nil
}

func FromBytes(hasher hash.Hash, byt []byte) (string, error) {

	code, err := digest(byt, hasher)
	if err != nil {
		return "", err
	}

	return code, nil
}
