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

func (that *Filesystem) NormalizePath(str string) string {

	return NormalizePath(str)
}

func (that *Filesystem) NormalizePathLocal(str string) string {

	return NormalizePathLocal(str)
}

func (that *Filesystem) NormalizePathOS(str string) string {

	return NormalizePathOS(str, that.OS)
}

func (that *Filesystem) Subdirectories(finalPath string, relativePath string) ([]string, error) {

	return Subdirectories(finalPath, relativePath, that.OS)
}
