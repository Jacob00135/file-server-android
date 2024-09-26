package models

type File struct {
	Name     string `json:"filename"`
	FileType bool   `json:"is_dir"`
	FileSize int64  `json:"size"`
}

type DbFile struct {
	ID         uint
	Path       string
	Permission uint
}
