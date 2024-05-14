package handlers

//Implement the following commands:
// 1. create project should implement the whole project with the GO4T stack

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/adam-fraga/ratel/errors"
	"github.com/adam-fraga/ratel/models/datatypes"
)

func InitProject() {
	//Capture user input
	var appName string
	io.WriteString(os.Stdout, "Choose a name for your application: ")
	fmt.Scanln(&appName)
	err := createProjectStructure(appName)
	if err != nil {
		fmt.Println(err)
	}
}

func createProjectStructure(appName string) error {
	fmt.Printf("Creating the project structure for application %s...", appName)
	jsonFolders, err := getProjectStructureFromJson()
	if err != nil {

		return &errors.Error{Type: "Project Structure Error", Msg: "Error parsing the folders.json file: " + err.Error()}
	}

	for _, folder := range jsonFolders {
		fmt.Println(folder)
		// createFolders(&folder)
	}
	// createFiles()
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

func getProjectStructureFromJson() ([]datatypes.Folder, error) {
	fmt.Println("Parsing the folders from Folders.json...")

	var folders []datatypes.Folder

	rootPath, _ := os.Getwd()
	folderJsonFilePath := filepath.Join(rootPath, "/data/folders.json")

	projectStructureJsonFile, err := os.Open(folderJsonFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, &errors.Error{Type: "Project Structure Error", Msg: "Error opening the folders.json file"}
	}
	defer projectStructureJsonFile.Close()

	err = json.NewDecoder(projectStructureJsonFile).Decode(&folders)
	if err != nil {
		fmt.Println(err)
		return nil, &errors.Error{Type: "Project Structure Error", Msg: "Error decoding the folders.json file"}
	}

	return folders, nil
}
