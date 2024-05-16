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
	fmt.Printf("Creating the project structure for application %s...\n", appName)
	jsonFolders, err := getProjectStructureFromJson()
	if err != nil {
		return &errors.Error{Type: "Project Structure Error", Msg: "Error parsing the folders.json file: " + err.Error()}
	}
	err = createFolders(jsonFolders, false, "")
	if err != nil {
		return &errors.Error{Type: "Project Structure Error", Msg: "Error creating the folders: " + err.Error()}
	}
	// createFiles()
	return nil
}

// Create the folders for the project
func createFolders(folders []datatypes.Folder, isSubfolder bool, parentFolder string) error {
	subfolders := []datatypes.Folder{}

	if len(subfolders) > 0 {
		err := createFolders(subfolders, true, parentFolder)
		if err != nil {
			return &errors.Error{Type: "Project Structure Error", Msg: "Error creating the subfolders: " + err.Error()}
		}
	}

	for _, folder := range folders {
		if isSubfolder {
			rootPath, _ := os.Getwd()
			folder.FolderName = filepath.Join(rootPath, folder.FolderName)
		}

		err := os.Mkdir(folder.FolderName, folder.Permissions)
		if err != nil {
			return &errors.Error{Type: "Project Structure Error", Msg: "Error creating the folder: " + err.Error()}
		}
		fmt.Printf("Folder %s with permission: %d successfully created \n", folder.FolderName, folder.Permissions)
		if len(folder.SubFolders) > 0 {
			for _, subfolder := range folder.SubFolders {
				subfolders = append(subfolders, subfolder)
			}
		}
		subfolders = nil
	}
	return nil
}

func createFiles() {
	populateFiles()
	fmt.Println("Creating the files...")
}

func populateFiles() {
	fmt.Print("Populating the files...")
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
