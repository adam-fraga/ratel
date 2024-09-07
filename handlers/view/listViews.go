package view

import (
	"fmt"
	"os"

	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

// ViewFiles struct to hold the view files
type ViewFiles struct {
	totalFiles uint16

	Templates  []View
	Partials   []View
	Layouts    []View
	Pages      []View
	Components []View
	Metadatas  []View
	Forms      []View
}

// ListViews function to list the views
func ListViews(viewType string) error {
	var customError = er.DevError{
		Type:       "Error",
		Origin:     "ListViews",
		FileOrigin: "listViews.go",
		Msg:        "",
	}

	var viewFiles ViewFiles

	switch viewType {
	case "":
		if err := viewFiles.setViewFiles(&viewFiles, "all"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "templates":
		if err := viewFiles.setViewFiles(&viewFiles, "templates"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "forms":
		if err := viewFiles.setViewFiles(&viewFiles, "forms"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)

	case "partials":
		if err := viewFiles.setViewFiles(&viewFiles, "partials"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "layouts":
		if err := viewFiles.setViewFiles(&viewFiles, "layouts"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "pages":
		if err := viewFiles.setViewFiles(&viewFiles, "pages"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "components":
		if err := viewFiles.setViewFiles(&viewFiles, "components"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "metadatas":
		if err := viewFiles.setViewFiles(&viewFiles, "metadatas"); err != nil {
			customError.Msg = fmt.Sprintf("Error getting the files to print to the stdout: " + err.Error())
			return &customError
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	default:
		if err := ut.RunCommand("./ratel", true, "view list --help"); err != nil {
			customError.Msg = fmt.Sprintf("Error running the command to show the help: " + err.Error())
			return &customError
		}
	}
	return nil
}

// printFilesToStdOut function to print the files to the stdout
func (*ViewFiles) printFilesToStdOut(viewFiles *ViewFiles) {
	viewFiles.totalFiles = 0

	viewFiles.Beautify(viewFiles.Templates, viewFiles)
	viewFiles.Beautify(viewFiles.Partials, viewFiles)
	viewFiles.Beautify(viewFiles.Layouts, viewFiles)
	viewFiles.Beautify(viewFiles.Pages, viewFiles)
	viewFiles.Beautify(viewFiles.Components, viewFiles)
	viewFiles.Beautify(viewFiles.Metadatas, viewFiles)
	viewFiles.Beautify(viewFiles.Forms, viewFiles)

	ut.PrintInfoMsg("\n   TOTAL")
	ut.PrintSuccessMsg(fmt.Sprintf("     %d\n", viewFiles.totalFiles))
}

func (*ViewFiles) setViewFiles(viewFiles *ViewFiles, fileType string) error {
	fileTypes := []string{"templates", "partials", "layouts", "pages", "components", "metadatas", "forms"}

	if fileType == "all" {
		if err := viewFiles.getAllView(viewFiles, fileTypes); err != nil {
			return &er.ClientError{Msg: fmt.Sprintf("Error getting all the files to show: " + err.Error())}
		}
	} else {
		if err := viewFiles.getViewFiles(viewFiles, fileType); err != nil {
			return &er.ClientError{Msg: fmt.Sprintf("Error getting the files to show: " + err.Error())}
		}
	}
	return nil
}

// getAllView function to get all the views from folders
func (*ViewFiles) getAllView(viewFiles *ViewFiles, fileTypes []string) error {
	for _, fileType := range fileTypes {
		path := "views/" + string(fileType) + "/"
		files, err := os.Open(path)
		defer files.Close()
		if err != nil {
			return &er.ClientError{Msg: fmt.Sprintf("Error opening the directory for the %s views", fileTypes)}
		}
		for {
			file, err := files.Readdir(1)
			if err != nil {
				break
			}
			switch fileType {
			case "templates":
				viewFiles.Templates = append(viewFiles.Templates, View{Name: file[0].Name(), Path: path, Type: fileType})
			case "partials":
				viewFiles.Partials = append(viewFiles.Partials, View{Name: file[0].Name(), Path: path, Type: fileType})
			case "layouts":
				viewFiles.Layouts = append(viewFiles.Layouts, View{Name: file[0].Name(), Path: path, Type: fileType})
			case "pages":
				viewFiles.Pages = append(viewFiles.Pages, View{Name: file[0].Name(), Path: path, Type: fileType})
			case "components":
				viewFiles.Components = append(viewFiles.Components, View{Name: file[0].Name(), Path: path, Type: fileType})
			case "metadatas":
				viewFiles.Metadatas = append(viewFiles.Metadatas, View{Name: file[0].Name(), Path: path, Type: fileType})
			case "forms":
				viewFiles.Forms = append(viewFiles.Forms, View{Name: file[0].Name(), Path: path, Type: fileType})
			}
		}
	}
	return nil
}

// getViewFiles function to get the view files
func (*ViewFiles) getViewFiles(viewFiles *ViewFiles, fileType string) error {
	path := "views/" + fileType + "/"

	files, err := os.Open(path)
	defer files.Close()
	if err != nil {
		return &er.ClientError{Msg: fmt.Sprintf("Error opening the directory for the %s views: "+err.Error(), fileType)}
	}
	for {
		file, err := files.Readdir(1)

		if err != nil {
			break
		}

		var fileElement = string(file[0].Name())
		switch fileType {
		case "templates":
			viewFiles.Templates = append(viewFiles.Templates, View{Name: fileElement, Path: path, Type: fileType})
		case "partials":
			viewFiles.Partials = append(viewFiles.Partials, View{Name: fileElement, Path: path, Type: fileType})
		case "layouts":
			viewFiles.Layouts = append(viewFiles.Layouts, View{Name: fileElement, Path: path, Type: fileType})
		case "pages":
			viewFiles.Pages = append(viewFiles.Pages, View{Name: fileElement, Path: path, Type: fileType})
		case "components":
			viewFiles.Components = append(viewFiles.Components, View{Name: fileElement, Path: path, Type: fileType})
		case "metadatas":
			viewFiles.Metadatas = append(viewFiles.Metadatas, View{Name: fileElement, Path: path, Type: fileType})
		case "forms":
			viewFiles.Forms = append(viewFiles.Forms, View{Name: fileElement, Path: path, Type: fileType})
		}
	}

	return nil
}

// Beautify function to beautify the view files before printing to the stdout
func (*ViewFiles) Beautify(viewFileList []View, viewFiles *ViewFiles) {
	var count uint8

	for _, file := range viewFileList {
		count++
		if count == 1 {
			ut.PrintInfoMsg(fmt.Sprintf("\n   ***%s***", file.Type))
		}
		viewFiles.totalFiles++
		ut.PrintSuccessMsg(fmt.Sprintf("     %s%s", file.Path, file.Name))
	}
}
