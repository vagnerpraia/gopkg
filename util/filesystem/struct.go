package gpfilesystem

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

func (util *Filesystem) NormalizePath(str string) string {

	return NormalizePath(str)
}

func (util *Filesystem) NormalizePathLocal(str string) string {

	return NormalizePathLocal(str)
}

func (util *Filesystem) NormalizePathOS(str string, os gpenum.OS) string {

	return NormalizePathOS(str, os)
}
