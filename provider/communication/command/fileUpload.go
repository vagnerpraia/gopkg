package gpcommand

type FileUpload struct {
	Local  string
	Remote string
	Update bool
}

func NewFileUpload(local string, remote string, update bool) *FileUpload {

	return &FileUpload{
		Local:  local,
		Remote: remote,
		Update: update,
	}
}
