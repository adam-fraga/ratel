/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package project contains handlers to execute the logic of the project system of ratel web framework
*/

package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	em "github.com/adam-fraga/ratel/embed"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/schollz/progressbar/v3"
)

// Project represent a project
type Project struct {
	ProjectName string
	Folders     []Folder
	Framework   string
	Reponame    string
	Files       []File
}

// Folder represent a folder in the project
type Folder struct {
	FolderName string   `json:"folderName"`
	SubFolders []Folder `json:"subFolders"`
	Files      []string `json:"files"`
}

// File represent a file in the project
type File struct {
	FileName     string `json:"fileName"`
	FileContent  string `json:"fileContent"`
	Extension    string `json:"extension"`
	FileLocation string `json:"fileLocation"`
}

// Init the project creation process
func CreateProject(appName string) error {
	err := CreateProjectStructure(appName)
	if err != nil {
		return &er.ProjectError{
			Msg:    "Error trying to create the project",
			Origin: "handlers/project/ => func CreateProject()",
			Err:    err,
		}
	}
	PopulateProjectFiles()
	return nil
}

// Create the project structure based on the data/projectStruct.json file
func CreateProjectStructure(appName string) error {
	projectStruct, err := getProjectStructFromJsonFile()
	ut.PrintInfoMsg(fmt.Sprintf("\n Creating project structure for the application %s...", appName))
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func CreateProjectStructure()",
		}
	}

	for _, folder := range projectStruct {
		CreateFolder(&folder)
		if err != nil {
			ut.PrintErrorMsg(err.Error())
			return &er.DevError{
				Msg:    err.Error(),
				Origin: "handlers/project/createProject.go => func CreateProjectStructure()",
			}
		}
	}

	ut.PrintSuccessMsg("\n Project structure successfully created\n")
	return nil
}

// Create a folder with the given permissions and create the files and subfolders inside it
func CreateFolder(folder *Folder) error {
	if folder.FolderName != "root" {
		err := os.Mkdir(folder.FolderName, os.FileMode(0755))
		if err != nil {
			ut.PrintErrorMsg(err.Error())
			return &er.DevError{
				Msg:    err.Error(),
				Origin: "handlers/project/createProject.go => func CreateFolder()",
			}
		}
	}

	if len(folder.Files) > 0 {
		for _, file := range folder.Files {
			err := CreateFile(file)
			if err != nil {
				ut.PrintErrorMsg(err.Error())
				return &er.DevError{
					Msg:    err.Error(),
					Origin: "handlers/project/createProject.go => func CreateFolder()",
				}
			}
		}
	}

	if len(folder.SubFolders) > 0 {
		for _, subFolder := range folder.SubFolders {
			err := CreateFolder(&subFolder)
			if err != nil {
				ut.PrintErrorMsg(err.Error())
				return &er.DevError{
					Msg:    err.Error(),
					Origin: "handlers/project/createProject.go => func CreateFolder()",
				}
			}
		}
	}
	return nil
}

// Create a file with the given permissions
func CreateFile(fileDestination string) error {
	file, err := os.Create(fileDestination)
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func CreateFile()",
		}
	}

	err = os.Chmod(fileDestination, os.FileMode(0644))
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func CreateFile()",
		}
	}

	defer file.Close()
	return nil
}

// Populate the project files with the data from the files/configs embeded folder using goroutines
func PopulateProjectFiles() error {

	files, err := GetFilesFromProject()

	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func PopulateProjectFiles()",
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(files["src"]))

	for i := range files["src"] {
		bar := ut.SetProgressBar(files["dst"][i])
		go processFile(&wg, files["src"], files["dst"], i, bar)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
	return nil
}

/*
Process the file copying the contents from the source file containing in the embedded folder configs
to the destination file in the project structure
*/
func processFile(wg *sync.WaitGroup, srcFiles []string, dstFiles []string, i int, pb *progressbar.ProgressBar) error {
	var embeddedConfigs = em.EmbeddedConfigs
	defer wg.Done()
	srcFileData, err := embeddedConfigs.ReadFile(srcFiles[i])
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func processFile()",
		}
	}

	dstFilePath := dstFiles[i]
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func processFile()",
		}
	}

	defer dstFile.Close()

	_, err = io.Copy(dstFile, bytes.NewReader(srcFileData)) // Copy file contents
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func processFile()",
		}

	}

	pb.Finish()
	fmt.Println()
	return nil
}

// Get the files from the project structure embedded file json and return a map with the source and destination files
func GetFilesFromProject() (map[string][]string, error) {
	var fileName string
	var srcFiles []string
	var dstFiles []string
	files := make(map[string][]string)

	dataConfigFilePath := "configs/"

	ut.PrintInfoMsg(" Populating the project files...\n")
	projectStruct, err := getProjectStructFromJsonFile()
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return nil, &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func GetFilesFromProject()",
		}
	}

	for _, folder := range projectStruct {
		for _, file := range folder.Files {
			fileName = filepath.Base(file)
			dstFiles = append(dstFiles, file)
			joinPath := filepath.Join(dataConfigFilePath, fileName)
			srcFiles = append(srcFiles, joinPath)
		}

		for _, subFolder := range folder.SubFolders {
			for _, file := range subFolder.Files {
				fileName = filepath.Base(file)
				dstFiles = append(dstFiles, file)
				joinPath := filepath.Join(dataConfigFilePath, fileName)
				srcFiles = append(srcFiles, joinPath)
			}
		}
	}

	files["src"] = srcFiles
	files["dst"] = dstFiles

	return files, nil
}

// Parse the project structure from the data/projectStruct.json file and return a slice of Folder structs
func getProjectStructFromJsonFile() ([]Folder, error) {
	var folders []Folder

	err := json.Unmarshal(em.EmbeddedProjectStruct, &folders)
	if err != nil {
		ut.PrintErrorMsg(err.Error())
		return nil, &er.DevError{
			Msg:    err.Error(),
			Origin: "handlers/project/createProject.go => func getProjectStructFromJsonFile()",
		}
	}

	return folders, nil
}
