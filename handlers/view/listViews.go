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
	e := er.ViewError{
		Origin: "File: handlers/view/listViews.go => Func: ListViews()",
		Msg:    "",
		Err:    nil,
	}

	var viewFiles ViewFiles

	switch viewType {
	case "":
		if err := viewFiles.setViewFiles(&viewFiles, "all"); err != nil {
			e.Msg = "Failed listing all views in the project, error: " + err.Error()
			e.Err = err
			return &e
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "templates":
		if err := viewFiles.setViewFiles(&viewFiles, "templates"); err != nil {
			e.Msg = "Failed listing templates views in the project, error: " + err.Error()
			e.Err = err
			return &e

		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "forms":
		if err := viewFiles.setViewFiles(&viewFiles, "forms"); err != nil {
			e.Msg = "Failed listing forms views in the project, error: " + err.Error()
			e.Err = err
			return &e

		}
		viewFiles.printFilesToStdOut(&viewFiles)

	case "partials":
		if err := viewFiles.setViewFiles(&viewFiles, "partials"); err != nil {
			e.Msg = "Failed listing partials views in the project, error: " + err.Error()
			e.Err = err
			return &e
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "layouts":
		if err := viewFiles.setViewFiles(&viewFiles, "layouts"); err != nil {
			e.Msg = "Failed listing layouts views in the project, error: " + err.Error()
			e.Err = err
			return &e

		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "pages":
		if err := viewFiles.setViewFiles(&viewFiles, "pages"); err != nil {
			e.Msg = "Failed listing pages views in the project, error: " + e.Error()
			e.Err = err
			return &e
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "components":
		if err := viewFiles.setViewFiles(&viewFiles, "components"); err != nil {
			e.Msg = "Failed listing components views in the project, error: " + err.Error()
			e.Err = err
			return &e
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	case "metadatas":
		if err := viewFiles.setViewFiles(&viewFiles, "metadatas"); err != nil {
			e.Msg = "Failed listing metadatas views in the project, error: " + err.Error()
			e.Err = err
			return &e
		}
		viewFiles.printFilesToStdOut(&viewFiles)
	default:
		if err := ut.RunCommand("ratel", true, "view list --help"); err != nil {
			e.Msg = "Failed running help command, error: " + err.Error()
			e.Err = err
			return &e
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

	ut.PrintInfoMsg("\n TOTAL")
	ut.PrintSuccessMsg(fmt.Sprintf(" %d\n", viewFiles.totalFiles))
}

func (*ViewFiles) setViewFiles(viewFiles *ViewFiles, fileType string) error {
	fileTypes := []string{"templates", "partials", "layouts", "pages", "components", "metadatas", "forms"}

	if fileType == "all" {
		if err := viewFiles.getAllView(viewFiles, fileTypes); err != nil {
			return &er.ViewError{
				Origin: "File: handlers/view/listViews.go => Func: setViewFiles()",
				Msg:    "Failed getting all views in the project, error: " + err.Error(),
				Err:    err,
			}
		}
	} else {
		if err := viewFiles.getViewFiles(viewFiles, fileType); err != nil {
			return &er.ViewError{
				Origin: "File: handlers/view/listViews.go => Func: setViewFiles()",
				Msg:    "Failed getting views files in the project, error: " + err.Error(),
				Err:    err,
			}
		}
	}
	return nil
}

// getAllView function to get all the views from folders
func (*ViewFiles) getAllView(viewFiles *ViewFiles, fileTypes []string) error {
	for _, fileType := range fileTypes {
		path := "src/views/" + string(fileType) + "/"
		files, err := os.Open(path)
		defer files.Close()
		if err != nil {
			return &er.ViewError{
				Origin: "File: handlers/view/listViews.go => Func: getAllviews()",
				Msg:    "Failed getting views, error opening view file the project, error: " + err.Error(),
				Err:    err,
			}
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
	path := "src/views/" + fileType + "/"

	files, err := os.Open(path)
	defer files.Close()
	if err != nil {
		return &er.ViewError{
			Origin: "File: handlers/view/listViews.go => Func: getViewFiles()",
			Msg:    "Failed getting views, error opening view file the project, error: " + err.Error(),
			Err:    err,
		}
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
			ut.PrintInfoMsg(fmt.Sprintf("\n  ***%s***", file.Type))
		}
		viewFiles.totalFiles++
		ut.PrintSuccessMsg(fmt.Sprintf("  %s%s", file.Path, file.Name))
	}
}
