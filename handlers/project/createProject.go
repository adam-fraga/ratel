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
func CreateProject(appName string) {
	CreateProjectStructure(appName)
	PopulateProjectFiles()
}

// Create the project structure based on the data/projectStruct.json file
func CreateProjectStructure(appName string) error {
	ut.PrintInfoMsg(fmt.Sprintf("   Creating the project structure folder for the application %s\n", appName))
	projectStruct, err := getProjectStructFromJsonFile()
	if err != nil {
		fmt.Println("Error parsing the json folders")
		fmt.Println(err.Error())
		return err
	}

	for _, folder := range projectStruct {
		CreateFolder(&folder)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	ut.PrintSuccessMsg("\n   Project structure successfully created\n")
	return nil
}

// Create a folder with the given permissions and create the files and subfolders inside it
func CreateFolder(folder *Folder) {
	ut.PrintInfoMsg(fmt.Sprintf("     %s Folder with permissions 755 successfuly created", folder.FolderName))

	if folder.FolderName != "root" {
		err := os.Mkdir(folder.FolderName, os.FileMode(0755))
		if err != nil {
			ut.PrintErrorMsg("Error creating folder " + folder.FolderName + ": " + err.Error())
		}
	}

	if len(folder.Files) > 0 {
		for _, file := range folder.Files {
			CreateFile(file)
		}
	}

	if len(folder.SubFolders) > 0 {
		for _, subFolder := range folder.SubFolders {
			CreateFolder(&subFolder)
		}
	}
}

// Create a file with the given permissions
func CreateFile(fileDestination string) {
	file, err := os.Create(fileDestination)
	if err != nil {
		ut.PrintErrorMsg("Error creating file")
	}

	err = os.Chmod(fileDestination, os.FileMode(0644))
	if err != nil {
		ut.PrintErrorMsg("Error setting right for file")
	}

	defer file.Close()
}

// Populate the project files with the data from the files/configs embeded folder using goroutines
func PopulateProjectFiles() {

	files := GetFilesFromProject()

	var wg sync.WaitGroup
	wg.Add(len(files["src"]))

	for i := range files["src"] {
		bar := ut.SetProgressBar(files["dst"][i])
		go processFile(&wg, files["src"], files["dst"], i, bar)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
}

/*
Process the file copying the contents from the source file containing in the embedded folder configs
to the destination file in the project structure
*/
func processFile(wg *sync.WaitGroup, srcFiles []string, dstFiles []string, i int, pb *progressbar.ProgressBar) {
	var embeddedConfigs = em.EmbeddedConfigs
	defer wg.Done()
	srcFileData, err := embeddedConfigs.ReadFile(srcFiles[i])
	if err != nil {
		ut.PrintErrorMsg("Error reading file from embedded source file " + srcFiles[i] + err.Error())
	}

	dstFilePath := dstFiles[i]
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		ut.PrintErrorMsg("Error creating file " + dstFiles[i] + err.Error())
	}

	defer dstFile.Close()

	_, err = io.Copy(dstFile, bytes.NewReader(srcFileData)) // Copy file contents
	if err != nil {
		ut.PrintErrorMsg("Error copiying file " + dstFiles[i] + err.Error())
	}

	pb.Finish()
	fmt.Println()
}

// Get the files from the project structure embedded file json and return a map with the source and destination files
func GetFilesFromProject() map[string][]string {
	var fileName string
	var srcFiles []string
	var dstFiles []string
	files := make(map[string][]string)

	dataConfigFilePath := "configs/"

	ut.PrintSuccessMsg("   Populating the project files...\n")
	projectStruct, err := getProjectStructFromJsonFile()
	if err != nil {
		ut.PrintErrorMsg("Error getting project structure from projectStruct.json: " + err.Error())
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

	return files
}

// Parse the project structure from the data/projectStruct.json file and return a slice of Folder structs
func getProjectStructFromJsonFile() ([]Folder, error) {
	var folders []Folder

	err := json.Unmarshal(em.EmbeddedProjectStruct, &folders)
	if err != nil {
		return nil, fmt.Errorf("error decoding embedded project struct JSON: %w", err)
	}

	return folders, nil
}
