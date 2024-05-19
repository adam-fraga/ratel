package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adam-fraga/ratel/errors"
	"github.com/adam-fraga/ratel/models/datatypes"
)

func InitProject(appName string) {

	createProjectStructure(appName)
}

func createProjectStructure(appName string) error {
	fmt.Printf("Creating the project structure for application %s...\n", appName)
	jsonFolders, err := parseJsonFolders()
	if err != nil {
		fmt.Println("Error parsing the json folders")
		fmt.Println(err.Error())
		return err
	}

	for _, folder := range jsonFolders {
		err := createFolder(&folder)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

func createFolder(folder *datatypes.Folder) error {
	fmt.Printf("Creating the folder %s with permissions 755...\n", folder.FolderName)

	if folder.FolderName != "root" {
		err := os.Mkdir(folder.FolderName, os.FileMode(0755))
		if err != nil {
			return &errors.Error{
				Type:       "Project Structure Error",
				Origin:     "createFolder()",
				FileOrigin: "handlers/project.go",
				Msg:        err.Error() + fmt.Sprintf("Error creating %s folder", folder.FolderName)}
		}
	}

	if len(folder.Files) > 0 {
		for _, file := range folder.Files {
			err := createFile(file)
			if err != nil {
				return &errors.Error{
					Type:       "Project Structure Error",
					Origin:     "createFolder()",
					FileOrigin: "handlers/project.go",
					Msg:        err.Error() + fmt.Sprintf("Error creating file %s in folder %s", file, folder.FolderName)}
			}
		}
	}

	if len(folder.SubFolders) > 0 {
		for _, subFolder := range folder.SubFolders {
			err := createFolder(&subFolder)
			if err != nil {
				return &errors.Error{
					Type:       "Project Structure Error",
					Origin:     "createFolder()",
					FileOrigin: "handlers/project.go",
					Msg:        err.Error() + fmt.Sprintf("Error creating subfolder %s in folder %s", subFolder.FolderName, folder.FolderName)}
			}
		}
	}

	return nil
}

func createFile(fileName string) error {
	fmt.Printf("Creating the file %s with permissions 755...\n", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return &errors.Error{
			Type:       "Project Structure Error",
			Origin:     "createFolder()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + fmt.Sprintf("Error creating %s file", fileName)}
	}
	defer file.Close()
	return nil
}

func populateFiles() {
	fmt.Println("Populating the files...")
}

func parseJsonFolders() ([]datatypes.Folder, error) {
	fmt.Println("Parsing the folders...")
	var folders []datatypes.Folder

	projectStructureJsonFilePath, err := filepath.Abs("/home/afraga/Projects/ratel/data/folders.json")

	if err != nil {
		return nil, &errors.Error{
			Type:       "Project Structure Error",
			Origin:     "parseJsonFolders()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + "Error getting the absolute path of the json file"}
	}

	jsonFile, err := os.Open(projectStructureJsonFilePath)
	if err != nil {
		return nil, &errors.Error{
			Type:       "Project Structure Error",
			Origin:     "parseJsonFolders()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error()}
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&folders)
	if err != nil {
		return nil, &errors.Error{
			Type:       "Project Structure Error",
			Origin:     "parseJsonFolders()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + "Error decoding the json file"}
	}
	return folders, nil
}
