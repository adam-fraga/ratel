package datatype

import (
	"io/fs"
)

type Project struct {
	ProjectName string
	Folders     []Folder
}

type Folder struct {
	FolderName  string
	SubFolders  []Folder
	Files       []File
	Permissions fs.FileMode
}

type File struct {
	FileName    string
	FileContent string
	Extension   string
	Permissions fs.FileMode
}
