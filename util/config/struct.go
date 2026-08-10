package gpconfig

import (
	gpenum "github.com/vagnerpraia/gopkg/enum"
)

type Filesystem struct {
	OS gpenum.OS
}

func NewFilesystem(location string, os gpenum.OS) *Filesystem {

	return &Filesystem{
		OS: os,
	}
}

func (util *Filesystem) NormalizePath(path string, output any) error {

	return Unmarshal(path, output)
}
