package datatypes

import (
	"io/fs"
)

type Project struct {
	ProjectName string
	Folders     []Folder
}

type Folder struct {
	FolderName  string      `json:"folderName"`
	SubFolders  []Folder    `json:"subFolders"`
	Files       []string    `json:"files"`
	Permissions fs.FileMode `json:"permissions"`
}

type File struct {
	FileName    string      `json:"fileName"`
	FileContent string      `json:"fileContent"`
	Extension   string      `json:"extension"`
	Permissions fs.FileMode `json:"permissions"`
}
