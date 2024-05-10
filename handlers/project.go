package handlers

//Implement the following commands:
// 1. create project should implement the whole project with the GO4T stack

import (
	"fmt"
	"os"

	"github.com/adam-fraga/ratel/datatype"
	"github.com/adam-fraga/ratel/errors"
	"github.com/joho/godotenv"
)

func InitProject() {
	fmt.Println("Creating a new project")
	createProjectStructure()
}

func createProjectStructure() {
	fmt.Println("Creating the project structure...")
	err, dotenv := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	appName := dotenv["APP_NAME"]

	folders := []datatype.Folder{
		{FolderName: appName, Permissions: 755},
		{FolderName: "cmd", Permissions: 755, SubFolders: []datatype.Folder{appName}},
		{FolderName: "static", Permissions: 755, SubFolders: []datatype.Folder{
			{FolderName: "style", Permissions: 755},
			{FolderName: "script", Permissions: 755}}},
		{FolderName: "src", Permissions: 755, SubFolders: []datatype.Folder{
			{FolderName: "style", Permissions: 755},
			{FolderName: "script", Permissions: 755}}},
		{FolderName: "handler", Permissions: 755},
		{FolderName: "views", Permissions: 755, SubFolders: []datatype.Folder{
			{FolderName: "views/layout", Permissions: 755},
			{FolderName: "views/metadatas", Permissions: 755},
			{FolderName: "views/templates", Permissions: 755},
			{FolderName: "views/partials", Permissions: 755},
			{FolderName: "views/components", Permissions: 755},
			{FolderName: "views/pages", Permissions: 755}}},
		{FolderName: "config", Permissions: 755},
		{FolderName: "script", Permissions: 755},
		{FolderName: "docs", Permissions: 755},

		{FolderName: "error", Permissions: 755},
		{FolderName: "db", Permissions: 755},
		{FolderName: "middleware", Permissions: 755},
		{FolderName: "model", Permissions: 755},
	}

	for _, folder := range folders {
		createFolders(&folder)
	}
	createFiles()
}

// Create the folders for the project
func createFolders(folder *datatype.Folder) error {
	fmt.Sprintf("Creating the folder %s...", folder.FolderName)
	err := os.Mkdir(folder.FolderName, folder.Permissions)

	if err != nil {
		return &errors.Error{Type: "Project Structure Error", Msg: fmt.Sprintf("Error creating %s folder", folder.FolderName)}
	}

	fmt.Sprintf("Folder %s successfully created", folder.FolderName)

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
