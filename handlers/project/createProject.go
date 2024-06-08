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
	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/schollz/progressbar/v3"
)

// Project represent a project
type Project struct {
	ProjectName string
	Folders     []Folder
}

// Folder represent a folder in the project
type Folder struct {
	FolderName string   `json:"folderName"`
	SubFolders []Folder `json:"subFolders"`
	Files      []string `json:"files"`
}

// File represent a file in the project
type File struct {
	FileName    string `json:"fileName"`
	FileContent string `json:"fileContent"`
	Extension   string `json:"extension"`
}

// Init the project creation process
func InitProject(appName string) {
	CreateProjectStructure(appName)
	PopulateProjectFiles()
}

// Create the project structure based on the data/projectStruct.json file
func CreateProjectStructure(appName string) error {
	ut.PrintInfoMsg(fmt.Sprintf("Creating the project structure for application %s", appName))
	projectStruct, err := getProjectStructFromJsonFile()
	if err != nil {
		fmt.Println("Error parsing the json folders")
		fmt.Println(err.Error())
		return err
	}

	for _, folder := range projectStruct {
		err := CreateFolder(&folder)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	ut.PrintSuccessMsg("\nProject structure successfully created\n")
	return nil
}

// Create a folder with the given permissions and create the files and subfolders inside it
func CreateFolder(folder *Folder) error {
	ut.PrintInfoMsg(fmt.Sprintf("%s Folder with permissions 755 successfuly created", folder.FolderName))

	if folder.FolderName != "root" {
		err := os.Mkdir(folder.FolderName, os.FileMode(0755))
		if err != nil {
			return &errors.DevError{
				Type:       "Project Structure Error",
				Origin:     "createFolder()",
				FileOrigin: "handlers/project.go",
				Msg:        err.Error() + fmt.Sprintf("Error creating %s folder", folder.FolderName)}
		}
	}

	if len(folder.Files) > 0 {
		for _, file := range folder.Files {
			err := CreateFile(file)
			if err != nil {
				return &errors.DevError{
					Type:       "Project Structure Error",
					Origin:     "createFolder()",
					FileOrigin: "handlers/project.go",
					Msg:        err.Error() + fmt.Sprintf("Error creating file %s in folder %s", file, folder.FolderName)}
			}
		}
	}

	if len(folder.SubFolders) > 0 {
		for _, subFolder := range folder.SubFolders {
			err := CreateFolder(&subFolder)
			if err != nil {
				return &errors.DevError{
					Type:       "Project Structure Error",
					Origin:     "createFolder()",
					FileOrigin: "handlers/project.go",
					Msg:        err.Error() + fmt.Sprintf("Error creating subfolder %s in folder %s", subFolder.FolderName, folder.FolderName)}
			}
		}
	}

	return nil
}

// Create a file with the given permissions
func CreateFile(fileDestination string) error {
	file, err := os.Create(fileDestination)
	if err != nil {
		return &errors.DevError{
			Type:       "Project Structure Error",
			Origin:     "createFolder()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + fmt.Sprintf("Error creating %s file\n", fileDestination)}
	}

	err2 := os.Chmod(fileDestination, os.FileMode(0777))

	if err2 != nil {
		return &errors.DevError{
			Type:       "Project Structure Error",
			Origin:     "createFolder()",
			FileOrigin: "handlers/project.go",
			Msg:        err2.Error() + fmt.Sprintf("Error changing permissions of %s file\n", fileDestination)}
	}

	defer file.Close()
	return nil
}

// Populate the project files with the data from the files/configs embeded folder using goroutines
func PopulateProjectFiles() {

	files, err := GetFilesFromProject()
	if err != nil {
		fmt.Println(err.Error())
	}

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
func processFile(wg *sync.WaitGroup, srcFiles []string, dstFiles []string, i int, pb *progressbar.ProgressBar) error {
	var embeddedConfigs = em.EmbeddedConfigs
	defer wg.Done()
	srcFileData, err := embeddedConfigs.ReadFile(srcFiles[i])
	if err != nil {
		return &errors.DevError{
			Type:       "Project Structure Error",
			Origin:     "processFile()",
			FileOrigin: "handlers/project.go",
			Msg:        fmt.Sprintf("Error reading file from embedded source file %s"+err.Error()+"\n", srcFiles[i])}
	}

	dstFilePath := dstFiles[i]
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return &errors.DevError{
			Type:       "Project Structure Error",
			Origin:     "processFile()",
			FileOrigin: "handlers/project.go",
			Msg:        fmt.Sprintf("Error creating file %s"+err.Error()+"\n", dstFiles[i])}
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, bytes.NewReader(srcFileData)) // Copy file contents
	if err != nil {
		return &errors.DevError{
			Type:       "Project Structure Error",
			Origin:     "processFile()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + fmt.Sprintf("Error copying file %s\n", dstFiles[i])}
	}

	pb.Finish()
	fmt.Println()
	ut.PrintSuccessMsg(dstFiles[i] + " successfully populated")

	return nil
}

// Get the files from the project structure embedded file json and return a map with the source and destination files
func GetFilesFromProject() (map[string][]string, error) {
	var fileName string
	var srcFiles []string
	var dstFiles []string
	files := make(map[string][]string)

	dataConfigFilePath := "configs/"

	ut.PrintInfoMsg("Populating the project files...")
	projectStruct, err := getProjectStructFromJsonFile()
	if err != nil {
		return nil, &errors.DevError{
			Type:       "Project Structure Error",
			Origin:     "getFilesFromProject()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + fmt.Sprintf("Error getting the project structure from json file\n")}
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
		return nil, fmt.Errorf("error decoding embedded project struct JSON: %w", err)
	}

	return folders, nil
}
