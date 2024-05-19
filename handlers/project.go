package handlers

//Implement the following commands:
// 1. create project should implement the whole project with the GO4T stack

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
		err := createFolder(&folder, "")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	// createFiles()
	return nil
}

func createFolder(folder *datatypes.Folder, parentFolder string) error {
	fmt.Printf("Creating the folder %s with permissions %d..\n", folder.FolderName, folder.Permissions)

	err := os.Mkdir(folder.FolderName, os.FileMode(0755))
	if err != nil {
		return &errors.Error{
			Type:       "Project Structure Error",
			Origin:     "createFolder()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + fmt.Sprintf("Error creating %s folder", folder.FolderName)}
	}

	if len(folder.SubFolders) > 0 {
		for _, subFolder := range folder.SubFolders {
			err := createFolder(&subFolder, folder.FolderName)
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

func createSubFolder(subFolder *datatypes.Folder, parentFolder string) error {
	fmt.Printf("Creating the subfolder %s inside parentFolder %s with permissions %d..", subFolder.FolderName, parentFolder, subFolder.Permissions)
	err := os.Mkdir(subFolder.FolderName, os.FileMode(0755))

	if err != nil {
		return &errors.Error{
			Type:       "Project Structure Error",
			Origin:     "createSubFolder()",
			FileOrigin: "handlers/project.go",
			Msg:        err.Error() + fmt.Sprintf("Error creating %s folder", subFolder.FolderName)}
	}

	return nil
}

func createFiles() {
	populateFiles()
	fmt.Println("Creating the files...")
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
