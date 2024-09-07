/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package project contains handlers to execute the logic of the project system of ratel web framework
*/

package project

import (
	"fmt"
	"os"

	s "strings"

	"github.com/adam-fraga/ratel/utils"
	"github.com/adam-fraga/ratel/internal/errors"
)

// Check first if the project is created (Presence of the project files)
// Ask user wich framework he want to use between None, Echo or Fiber

// Init the project creation process
func InitProject(reponame string, framework string) {
  p := Project{}

  if reponame == "" || !s.HasPrefix(reponame, "github.com/"){
    utils.PrintErrorMsg("Repo name is not well formatted: \"github.com/your-name/repo\"")
    return
  }

  p.Reponame = reponame

  if framework == "" {
    utils.RunCommandWithOutput("go", fmt.Sprintf("mod init %s", reponame))
    utils.PrintSuccessMsg(fmt.Sprintf("Successfully initialized repo: \"%s\" !", reponame))
    return
  } else if framework == "Echo" || framework == "Fiber"{
    p.Framework = framework
    utils.RunCommandWithOutput("go", fmt.Sprintf("mod init %s", p.Reponame))
    utils.PrintSuccessMsg(fmt.Sprintf("Successfully initialized repo: %s with %s framework !", p.Reponame, p.Framework))
    getFrameworkFromGoPackage(&p)
  }
}

func getFrameworkFromGoPackage(p *Project){
  if p.Framework == "Echo" {
    utils.PrintInfoMsg("Get echo framework from go package...")
    utils.RunCommandWithOutput("go", "get github.com/labstack/echo/v4")
  } else if p.Framework == "Fiber" {
    utils.PrintInfoMsg("Get echo framework from go package...")
    utils.RunCommandWithOutput("go", "get github.com/gofiber/fiber/v2")
  }
  err := writeFiles(p)
  if err != nil {
    fmt.Println(err.Error())
  }
}


func writeFiles(p *Project) error {
  files := []File{
    {FileName : "main",  Extension: ".go", FileLocation : "./cmd/"},
    {FileName : "index", Extension: ".go", FileLocation : "./handlers/"},
    {FileName : "db", Extension: ".go", FileLocation : "./db/"},
    {FileName : "router", Extension: ".go", FileLocation : "./routers/"},
    {FileName : "user", Extension: ".go", FileLocation : "./models/"},
    {FileName : "errors", Extension: ".go", FileLocation : "./error/"},
    {FileName : "auth-middleware", FileContent : "", Extension: ".go", FileLocation : "./middlewares/"},
    {FileName : "utils", Extension: ".go", FileLocation : "./utils/"},
    {FileName : "header", Extension: ".templ", FileLocation : "./views/partials/"},
    {FileName : "footer", Extension: ".templ", FileLocation : "./views/partials/"},
    {FileName : "base", Extension: ".templ", FileLocation : "./views/layouts/"},
    {FileName : "index",  Extension: ".templ", FileLocation : "./views/pages/"},
    {FileName : "index.test", Extension: ".go", FileLocation : "./test/"},
  }

  p.Files  = files

  for _, f := range p.Files {
    go writeFile(f, p)
  }


  return nil
}

func writeFile(f File, p *Project) error {

  fileLocation := fmt.Sprintf("%s%s%s", f.FileName, f.Extension, f.FileLocation)
  fileName := fmt.Sprintf("%s%s", f.FileName, f.Extension)

	file, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return &errors.ClientError{
			Msg: err.Error() + fmt.Sprintf("Error trying to open file \"%s\". You need to be at the root of your project after creating it with \"ratel project create\".\n", fileLocation),
		}
	}
	defer file.Close()

	var content string

  switch fileName {
  case "main.go":
    if p.Framework == "Fiber" {
    content = fmt.Sprintf(`
    package main

    import (
      "github.com/gofiber/fiber/v2"
      h "%s/handlers" // Change path to match your project path and uncomment
      r "%s/router"   // Change path to match your project path and uncomment

    )

    func main() {
      app := fiber.New()

      app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Ratel!")
      })

      app.Listen(":8080")
    }
    `, p.Reponame, p.Reponame)
      } else if p.Framework == "Echo" {
        content = fmt.Sprintf(`
    package main

    import (
      "net/http"
      "github.com/labstack/echo/v4"
      "github.com/labstack/echo/v4/middleware"
     h "%s/handlers"
     r "%s/router" 

    )

    func main() {
      e := echo.New()

      e.Use(middleware.Logger())
      e.Use(middleware.Recover())

      e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, Ratel!")
      })

      e.Logger.Fatal(e.Start(":8080"))
    }
    `, p.Reponame, p.Reponame)
    } 
    break
  case "index.go":
    content = fmt.Sprintf(`
      package handlers

import (
	"net/http"
	 t "github.com/a-h/templ"
	 pages "%s/views/pages" #Change path to match your project and uncomment
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	homeData := make(map[string]interface{})
	homeData["name"] = "John"
	// homePage := pages.Index(homeData)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	// t.Handler(homePage).ServeHTTP(w, r)
}`, p.Reponame)
      break
  case "index.test.go":
    content = "package test;"
    break
  case "db.go":
    content = "package db;"
      break
  case "router.go":
    content = "package routers;"
      break
  case "user.go":
    content = "package models;"
      break
  case "error.go":
    content = "package error;"
      break
  case "utils.go":
    content = "package utils;"
      break
  case "auth-middleware.go":
    content = "package middlewares;"
      break
  case "header.templ":
    break
  case "footer.templ":
    break
  case "base.templ":
    break
  case "index.templ":
    break
   }
  _, err = file.WriteString(content)
  if err != nil {
    return &errors.ClientError{
      Msg: err.Error() + fmt.Sprintf("Error trying to write in file \"./cmd/main.go\". Ensure you're at the project root.\n"),
    }
  }

	return nil
}

