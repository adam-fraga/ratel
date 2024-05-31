package views

import (
	"fmt"
	"os"

	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

func ListViews(viewType string) error {
	var customError = er.DevError{
		Type:       "Error",
		Origin:     "ListViews",
		FileOrigin: "listViews.go",
		Msg:        "",
	}

	switch viewType {
	case "":
		filesToPrint, err := getFilesToPrintToStdOut("all")
		if err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		printFilesToStdOut(filesToPrint)
	case "templates":
		filesToPrint, err := getFilesToPrintToStdOut("templates")
		if err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		printFilesToStdOut(filesToPrint)
	case "partials":
		filesToPrint, err := getFilesToPrintToStdOut("partials")
		if err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		printFilesToStdOut(filesToPrint)
	case "layouts":
		filesToPrint, err := getFilesToPrintToStdOut("layouts")
		if err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		printFilesToStdOut(filesToPrint)
	case "pages":
		filesToPrint, err := getFilesToPrintToStdOut("pages")
		if err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		printFilesToStdOut(filesToPrint)
	case "components":
		filesToPrint, err := getFilesToPrintToStdOut("components")
		if err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		printFilesToStdOut(filesToPrint)
	case "metadatas":
		filesToPrint, err := getFilesToPrintToStdOut("metadatas")
		if err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		printFilesToStdOut(filesToPrint)
	default:
		ut.PrintErrorMsg("Choose a valid view type\n")
	}
	return nil
}

func getFilesToPrintToStdOut(fileType string) ([]string, error) {

	fileTypes := []string{"templates", "partials", "layouts", "pages", "components", "metadatas"}

	path := "views/" + fileType + "/"

	var fileList []string

	if fileType == "all" {
		for _, ft := range fileTypes {
			path := "views/" + string(ft) + "/"
			files, err := os.Open(path)
			defer files.Close()
			if err != nil {
				return nil, &er.ClientError{Msg: fmt.Sprintf("Error opening the directory for the %s views", fileTypes)}
			}
			for {
				file, err := files.Readdir(1)
				if err != nil {
					break
				}
				var fileElement = string(file[0].Name())
				fileList = append(fileList, fileElement)
			}
		}
	} else {
		files, err := os.Open(path)
		defer files.Close()
		if err != nil {
			return nil, &er.ClientError{Msg: fmt.Sprintf("Error opening the directory for the %s views", fileType)}
		}
		for {
			file, err := files.Readdir(1)
			if err != nil {
				break
			}
			var fileElement = string(file[0].Name())
			fileList = append(fileList, fileElement)
		}
	}

	return fileList, nil
}

func printFilesToStdOut(files []string) {
	for key, file := range files {
		fmt.Printf("%d - %s\n", key, file)
	}
}
