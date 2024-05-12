package handlers

//Implement the following commands:
// 1. create project should implement the whole project with the GO4T stack

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adam-fraga/ratel/errors"
	"github.com/adam-fraga/ratel/models/datatypes"
)

func InitProject(appName string) {
	fmt.Println("Creating a new project")
	createProjectStructure(appName)
}

func createProjectStructure(appName string) error {
	fmt.Printf("Creating the project structure for application %s...", appName)
	jsonFolders, err := parseJsonFolders()
	if err != nil {
		return &errors.Error{Type: "Project Structure Error", Msg: "Error parsing the folders.json file"}
	}

	for _, folder := range jsonFolders {
		createFolders(&folder)
	}
	createFiles()
	return nil
}

// Create the folders for the project
func createFolders(folder *datatypes.Folder) error {
	fmt.Printf("Creating the folder %s...", folder.FolderName)

	err := os.Mkdir(folder.FolderName, folder.Permissions)
	if err != nil {
		return &errors.Error{Type: "Project Structure Error", Msg: fmt.Sprintf("Error creating %s folder", folder.FolderName)}
	}

	fmt.Printf("Folder %s successfully created", folder.FolderName)
	if len(folder.SubFolders) > 0 {
		for _, subFolder := range folder.SubFolders {
			createFolders(&subFolder)
		}
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

	jsonFile, err := os.Open("folders.json")
	if err != nil {
		return nil, &errors.Error{Type: "Project Structure Error", Msg: "Error opening the folders.json file"}
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&folders)
	if err != nil {
		return nil, &errors.Error{Type: "Project Structure Error", Msg: "Error decoding the folders.json file"}
	}
	return folders, nil
}
