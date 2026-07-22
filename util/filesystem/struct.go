package gpfilesystem

type Filesystem struct {
	OS OS
}

func NewFilesystem(location string, os OS) *Filesystem {

	return &Filesystem{
		OS: os,
	}
}

func (util *Filesystem) NormalizePath(str string, os OS) string {

	return NormalizePath(str, os)
}
