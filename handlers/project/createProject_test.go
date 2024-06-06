package project

import (
	"testing"
)

func TestCreateProjectStructure(t *testing.T) {
	err := CreateProjectStructure("test")
	if err != nil {
		t.Errorf("Error creating project structure: %s", err.Error())
	}
}

func TestCreateFolder(t *testing.T) {
	ProjectStructure := Folder{
		FolderName: "tmp/test",
		SubFolders: []Folder{},
		Files:      []string{},
	}

	ProjectStructure.SubFolders = append(ProjectStructure.SubFolders, Folder{
		FolderName: "tmp/test/subtest",
		SubFolders: []Folder{},
		Files:      []string{},
	})

	err := CreateFolder(&ProjectStructure)
	if err != nil {
		t.Errorf("Error creating folder: %s", err.Error())
	}

}

func TestCreateFile(t *testing.T) {
	ProjectStructure := Folder{
		FolderName: "tmp/test",
		SubFolders: []Folder{},
		Files:      []string{},
	}

	ProjectStructure.Files = append(ProjectStructure.Files, "testFile")
	ProjectStructure.Files = append(ProjectStructure.Files, "testFil2")

	for _, file := range ProjectStructure.Files {
		err := CreateFile(file)
		if err != nil {
			t.Errorf("Error creating file: %s", err.Error())
		}
	}

}
