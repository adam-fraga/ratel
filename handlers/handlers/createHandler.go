package handler

import (
	"fmt"
	"os"
	"path"

	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

type Handler struct {
	Name string
	Path string
}

func (*Handler) New() *Handler {
	return &Handler{}
}

// CreateView Create a file view of type (Component, Page, Layout)
func (*Handler) Create(handlers []string) error {

	var h Handler

	if len(handlers) > 1 {
		ut.PrintInfoMsg("\n   Creating multiple handlers\n")

		for _, handler := range handlers {
			ut.PrintSuccessMsg(fmt.Sprintf("     %s", handler))
		}

		ut.PrintWarningMsg("\n   Confirm you to create the followings handlers (y/n):")
		var response string

		fmt.Scanln(&response)

		if response == "n" {
			h.Create(handlers)
		}
	}

	for _, handler := range handlers {
		h.Name = handler
		if h.Name == "" {
			return &errors.ClientError{Msg: fmt.Sprintf("%s name cannot be empty", h.Name)}
		}

		if err := h.CreateFile(h); err != nil {
			return &errors.DevError{Msg: fmt.Sprintf("Error creating %s file :" + err.Error())}
		}
	}

	return nil

}

func (*Handler) CreateFile(h Handler) error {

	h.Path = "handlers/"

	file, err := os.Create(path.Join(h.Path, h.Name+".go"))
	defer file.Close()

	if err != nil {
		return &errors.DevError{
			Type:       "Creation view file error",
			Origin:     "createViewFile()",
			FileOrigin: "handlers/middlewares/createMiddleware.go",
			Msg:        err.Error() + fmt.Sprintf("Error creating %v file\n", file)}
	}

	if err := os.Chmod(h.Path+h.Name+".go", os.FileMode(0755)); err != nil {
		return &errors.DevError{
			Type:       "Creation view file error",
			Origin:     "createViewFile()",
			FileOrigin: "handlers/views/createView.go",
			Msg:        err.Error() + fmt.Sprintf("Error changing permissions for %v file\n", file)}
	}

	ut.PrintSuccessMsg(fmt.Sprintf("     %s%s.go successfuly created", h.Path, h.Name))
	return nil
}
