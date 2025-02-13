/*
Copyright © 2024 Adm FRG adam.fraga@admtechlabs.com
Package project contains handlers to execute the logic of the project system of ratel web framework
*/

package project

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	s "strings"

	"github.com/adam-fraga/ratel/internal/errors"
	"github.com/adam-fraga/ratel/utils"
)

// Check first if the project is created (Presence of the project files)
// Ask user wich framework he want to use between None, Echo or Fiber

// Init the project creation process
func InitProject(reponame string, framework string, dbProvider string) error {

	p := Project{}
	p.Framework = strings.ToLower(framework)
	p.DBProvider = strings.ToLower(dbProvider)

	files := []File{
		{FileName: "main", Extension: ".go", FileLocation: "."},
		{FileName: "index", Extension: ".go", FileLocation: "./src/handlers/"},
		{FileName: "db", Extension: ".go", FileLocation: "./src/db/"},
		{FileName: "user", Extension: ".go", FileLocation: "./src/models/"},
		{FileName: "errors", Extension: ".go", FileLocation: "./src/errors/"},
		{FileName: "utils", Extension: ".go", FileLocation: "./src/utils/"},
		{FileName: "header", Extension: ".templ", FileLocation: "./src/views/partials/"},
		{FileName: "footer", Extension: ".templ", FileLocation: "./src/views/partials/"},
		{FileName: "base", Extension: ".templ", FileLocation: "./src/views/layouts/"},
		{FileName: "index", Extension: ".templ", FileLocation: "./src/views/pages/"},
		{FileName: "index.test", Extension: ".go", FileLocation: "./src/test/"},
	}

	p.Files = files

	if reponame == "" || !s.HasPrefix(reponame, "github.com/") {
		return &errors.ProjectError{
			Origin: "File: handlers/project/initProject.go => Func: initProject()",
			Msg:    "Repo name is not well formatted: \"github.com/your-name/repo\"",
			Err:    nil,
		}
	}
	p.Reponame = reponame

	initGoModule(p.Reponame)

	if p.DBProvider == "postgres" { //Rename to match pq in the package dependencies name
		p.DBProvider = "pq"
	}

	utils.PrintCyanInfoMsg(" 🛠️ Preparing project files...\n")
	err := writeFiles(&p)
	if err != nil {
		return &errors.ProjectError{
			Origin: "File: handlers/project/initProject.go => Func: initProject()",
			Msg:    "Failed initialize project, error writing file: " + err.Error(),
			Err:    err,
		}
	}

	if err := getProjectDependencies(&p); err != nil {
		return &errors.ProjectError{
			Origin: "File: handlers/project/initProject.go => Func: initProject()",
			Msg:    "Failed initialize project, error getting packages: " + err.Error(),
			Err:    err,
		}
	}
	return nil
}

// Parse each file and write them concurrently
func writeFiles(p *Project) error {

	var wg sync.WaitGroup

	errChan := make(chan error, len(p.Files))

	defer close(errChan)

	for _, file := range p.Files {
		wg.Add(1)
		go writeFile(file, p, &wg, errChan)
	}

	for _, file := range p.Files {
		if err := <-errChan; err != nil {
			fmt.Printf(" Error writing file %s%s: %s", file.FileName, file.Extension, err.Error())
			os.Exit(1)
		}
	}

	wg.Wait()
	return nil
}

// Contains logic that Write files content  in the project
func writeFile(f File, p *Project, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	fileLocation := fmt.Sprintf("%s%s%s", f.FileLocation, f.FileName, f.Extension)
	fileName := fmt.Sprintf("%s%s", f.FileName, f.Extension)

	file, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		utils.PrintErrorMsg(err.Error())
		errChan <- &errors.ProjectError{
			Origin: "File: handlers/project/initProject.go => Func: writeFile()",
			Msg:    "Failed initialize project, error writing file " + fileName,
			Err:    nil,
		}
	} else {
		errChan <- nil
	}

	utils.PrintInfoMsg(fmt.Sprintf("📝 Writing file %s in location => %s", fileName, fileLocation))

	defer file.Close()

	var content string

	switch fileName {

	case "main.go":
		if p.Framework == "Fiber" {
			content = fmt.Sprintf(`package main

    import (
      "github.com/gofiber/fiber/v2"
      "%s/src/handlers"
    )

    func main() {
      app := fiber.New()

      app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Ratel!")
      })

      app.Listen(":8080")
    }
    `, p.Reponame)
		} else if p.Framework == "Echo" {
			content = fmt.Sprintf(`package main

    import (
      "net/http"
      "github.com/labstack/echo/v4"
      "github.com/labstack/echo/v4/middleware"
     "%s/src/handlers"
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
    `, p.Reponame)
		} else {
			content = fmt.Sprintf(`package main

import (
  "fmt"
	"net/http"
	"%s/src/handlers" 
	"%s/src/views/pages"
)

        func main() {
	mux := http.NewServeMux()
	//Serve Static files within static directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//Basic Handlers Example
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.Handle("/about", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html") // Ensure it's interpreted as HTML
		data := map[string]interface{}{
			"title": "About Page", // Adjust based on your template needs
		}
		err := pages.About(data).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}`, p.Reponame, p.Reponame)
		}
		break

	case "index.go":
		content = fmt.Sprintf(`package handlers

import (
	"net/http"
	 "github.com/a-h/templ"
	 "%s/src/views/pages"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	homePage := pages.Index()
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	templ.Handler(homePage).ServeHTTP(w, r)
      }`, p.Reponame)
		break

	case "index.test.go":
		content = "package test;"
		break

	case "db.go":
		content = "package db;"
		break

	case "utils.go":
		content = "//TODO WRITE UTILS"
		break

	case "header.templ":
		content = `package partials

templ Header() {
<header>
  <nav>
    <ul>
      <li><a href="/">Home</a></li>
      <li><a href="/about">About</a></li>
      <li><a href="/contact">Contact</a></li>
    </ul>
  </nav>
</header>
}`
		break

	case "footer.templ":
		content = `package partials

templ footer(data map[string]interface{}) {
<footer>
  <p>© 2017 Company, Inc.</p>
</footer>
}`
		break

	case "base.templ":
		content = fmt.Sprintf(`package layouts

import (
 "%s/src/views/metadatas"
 "%s/src/views/partials"
 "github.com/a-h/templ"
)

var (
Head = metadatas.Head
Header = partials.Header
Footer = partials.Footer
)

templ Base(Page t.Component) {
<!DOCTYPE html>
<html>
@Head()
@Header()

<body>
  @Page
</body>

@Footer()

</html>
}`, p.Reponame, p.Reponame)
		break

	case "index.templ":
		content = fmt.Sprintf(`package pages

templ Index() {
<main>
  <h1>Welcome</h1>
  <p>This is a simple example of a Go web app</p>
</main>
}

templ About(data map[string]interface{}) {
<main>
  <h1>About</h1>
  <p>This is a simple example of a Go web app</p>
</main>
}

templ Contact(data map[string]interface{}) {
<main>
  <h1>Contact</h1>
  <p>This is a simple example of a Go web app</p>
</main>
}`)
		break
	}

	_, err = file.WriteString(content)
	if err != nil {
		utils.PrintErrorMsg(err.Error())
		errChan <- &errors.ProjectError{
			Origin: "File: handlers/project/initProject.go => Func: writeFile()",
			Msg:    "Failed initialize project, error writing content in file " + fileName,
			Err:    nil,
		}

	} else {
		errChan <- nil
	}
}

// Get the project dependencies with go get command
func getProjectDependencies(p *Project) error {

	utils.PrintCyanInfoMsg("\n 📦 Fetching and installing dependencies...\n")

	customDependenciesCmd := [][]string{
		{"go", "get", "github.com/gofiber/fiber/v3"},
		{"go", "get", "github.com/labstack/echo/v4"},
		{"go", "get", "github.com/lib/pq"},
		{"go", "get", "go.mongodb.org/mongo-driver/v2/mongo"},
	}

	defaultDependenciesCmd := [][]string{
		{"go", "get", "github.com/mattn/go-sqlite3"},
		{"go", "get", "github.com/a-h/templ"},
		{"go", "get", "github.com/joho/godotenv"},
		{"mv", ".main.go", "main.go"},
		{"templ", "generate"},
		{"air"},
	}

	commands := filterDependenciesCommands(p, customDependenciesCmd)
	cmds := make([][]string, len(defaultDependenciesCmd)+len(commands))
	copy(cmds, commands)
	copy(cmds[len(commands):], defaultDependenciesCmd)

	for _, cmdArgs := range cmds {
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return &errors.ProjectError{
				Origin: "File: handlers/project/initProject.go => Func: getProjectPackages()",
				Msg:    "Failed running command to get project packages: " + strings.Join(cmdArgs, ", error: "+err.Error()),
				Err:    err,
			}
		}
		// Add messages based on the executed command
		switch strings.Join(cmdArgs, " ") {
		case "go get github.com/a-h/templ":
			utils.PrintSuccessMsg("✅ go: added github.com/a-h/templ")
		case "go get github.com/joho/godotenv":
			utils.PrintSuccessMsg("✅ go: added github.com/joho/godotenv latest")
		case "templ generate":
			utils.PrintCyanInfoMsg("⚡ Generating Templ components...")
			utils.PrintSuccessMsg("✅ Templ components generated successfully!")
		case "go get github.com/gofiber/fiber/v3":
			utils.PrintSuccessMsg("✅ go: added github.com/gofiber/fiber/v3 successfully!")
		case "go get github.com/labstack/echo/v4":
			utils.PrintSuccessMsg("✅ go: added github.com/echo/v4 successfully!")
		case "go get github.com/lib/pq":
			utils.PrintSuccessMsg("✅ go: added github.com/lib/pq successfully!")
		case "go get go.mongodb.org/mongo-driver/v2/mongo":
			utils.PrintSuccessMsg("✅ go: added github.com/mongodrive/v2/mongo successfully!")
		}
	}

	utils.PrintSuccessMsg("\n 🎉 Dependencies installed successfully. Your project is ready to go !\n")
	return nil
}

// Init go project with the go mod command
func initGoModule(reponame string) error {
	fmt.Println("")
	utils.PrintCyanInfoMsg("\n ⚙️ Initializing go module...\n")

	initCmd := exec.Command("go", "mod", "init", reponame)
	initCmd.Stdout = os.Stdout
	initCmd.Stderr = os.Stderr

	if err := initCmd.Run(); err != nil {
		return &errors.ProjectError{
			Origin: "File: handlers/project/initProject.go => Func: getProjectPackages()",
			Msg:    "Failed trying to init go module, error: " + err.Error(),
			Err:    err,
		}
	}

	utils.PrintSuccessMsg(fmt.Sprintf("\n ✅ Successfully initialized repo %s\n", reponame))
	return nil
}

// Filter commands to extract choosen Framework and db provider
func filterDependenciesCommands(p *Project, commands [][]string) [][]string {
	var cmds [][]string

	for _, cmdArgs := range commands {
		if strings.Contains(cmdArgs[2], p.Framework) {
			cmds = append(cmds, cmdArgs)
		}
		if strings.Contains(cmdArgs[2], p.DBProvider) {
			cmds = append(cmds, cmdArgs)
		}
	}
	return cmds
}
