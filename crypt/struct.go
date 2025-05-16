package gpcrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"hash/fnv"
)

type Crypt struct {
	hasher hash.Hash
}

func NewCrypt(hasher hash.Hash) *Crypt {

	return &Crypt{
		hasher: hasher,
	}
}

func NewCryptFNV1a32() *Crypt {

	return &Crypt{
		hasher: fnv.New32a(),
	}
}

func NewCryptFNV1a64() *Crypt {

	return &Crypt{
		hasher: fnv.New64a(),
	}
}

func NewCryptFNV1a128() *Crypt {

	return &Crypt{
		hasher: fnv.New128a(),
	}
}

func NewCryptMD5() *Crypt {

	return &Crypt{
		hasher: md5.New(),
	}
}

func NewCryptSHA1() *Crypt {

	return &Crypt{
		hasher: sha1.New(),
	}
}

func NewCryptSHA256() *Crypt {

	return &Crypt{
		hasher: sha256.New(),
	}
}

func (crypt *Crypt) FromInterface(i interface{}) (string, error) {

	return FromInterface(crypt.hasher, i)
}

func (crypt *Crypt) FromInterfaces(i ...interface{}) (string, error) {

	return FromInterfaces(crypt.hasher, i)
}

func (crypt *Crypt) FromString(str string) (string, error) {

	return FromString(crypt.hasher, str)
}

func (crypt *Crypt) FromBytes(byt []byte) (string, error) {

	return FromBytes(crypt.hasher, byt)
}
