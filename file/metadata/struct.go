package gpmetadata

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"hash/fnv"
	"io"
	"os"
)

type Metadata struct {
	Path     string
	Size     int64
	Checksum string
}

func NewMetadata(path string, hasher hash.Hash) (*Metadata, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(hasher, file); err != nil {
		return nil, err
	}

	return &Metadata{
		Path:     path,
		Size:     info.Size(),
		Checksum: hex.EncodeToString(hasher.Sum(nil)),
	}, nil
}

func NewCryptFNV1a32(path string) (*Metadata, error) {

	return NewMetadata(path, fnv.New32a())
}

func NewCryptFNV1a64(path string) (*Metadata, error) {

	return NewMetadata(path, fnv.New64a())
}

func NewCryptFNV1a128(path string) (*Metadata, error) {

	return NewMetadata(path, fnv.New128a())
}

func NewMetadataMD5(path string) (*Metadata, error) {

	return NewMetadata(path, md5.New())
}

func NewMetadataSHA1(path string) (*Metadata, error) {

	return NewMetadata(path, sha1.New())
}

func NewMetadataSHA256(path string) (*Metadata, error) {

	return NewMetadata(path, sha256.New())
}
