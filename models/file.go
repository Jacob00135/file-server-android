package models

type File struct {
	Name     string `json:"filename"`
	FileType bool   `json:"is_dir"`
	FileSize int64  `json:"size"`
}

type DbFile struct {
	ID         uint   `json:"id"`
	Path       string `json:"path"`
	Permission uint   `json:"permission"`
}
