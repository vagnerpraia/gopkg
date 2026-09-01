package gpcommand

type FileUpload struct {
	Local  string
	Remote string
}

func NewFileUpload(local string, remote string) *FileUpload {

	return &FileUpload{
		Local:  local,
		Remote: remote,
	}
}
