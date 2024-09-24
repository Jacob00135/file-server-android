package models

type File struct {
	ID         int    `json:"id"`
	Path       string `json:"path"`
	Permission int    `json:"permission"` // 例如: "read,write"
	FileType   string `json:"type"`       // 例如: "file" 或 "directory"
	FileSize   int    `json:"size"`
}

type DbFile struct {
	ID         int
	Path       string
	Permission int
}
