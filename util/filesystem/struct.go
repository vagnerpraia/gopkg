package gpfilesystem

type Filesystem struct {
	OS OS
}

func NewFilesystem(location string, os OS) *Filesystem {

	return &Filesystem{
		OS: os,
	}
}

func (util *Filesystem) NormalizePathLocal(str string) string {

	return NormalizePathLocal(str)
}

func (util *Filesystem) NormalizePathOS(str string, os OS) string {

	return NormalizePathOS(str, os)
}
