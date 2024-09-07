/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package project contains handlers to execute the logic of the project system of ratel web framework
*/

package project

import (
	"fmt"
	"os"
	"sync"

	s "strings"

	"github.com/adam-fraga/ratel/internal/errors"
	"github.com/adam-fraga/ratel/utils"
)

// Check first if the project is created (Presence of the project files)
// Ask user wich framework he want to use between None, Echo or Fiber

// Init the project creation process
func InitProject(reponame string, framework string) {

  p := Project{}

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

  if reponame == "" || !s.HasPrefix(reponame, "github.com/"){
    utils.PrintErrorMsg("Repo name is not well formatted: \"github.com/your-name/repo\"")
    return
  }

  p.Reponame = reponame

  if framework == "" {
    utils.RunCommand("go",false, fmt.Sprintf("mod init %s", reponame))
    utils.PrintSuccessMsg(fmt.Sprintf("\n Successfully initialized repo %s\n", reponame))
  } else if framework == "Echo" || framework == "Fiber"{
    p.Framework = framework
    utils.RunCommand("go", false, fmt.Sprintf("mod init %s", p.Reponame))
    utils.PrintSuccessMsg(fmt.Sprintf("\n Successfully initialized repo: %s with %s framework\n", p.Reponame, p.Framework))
    getFrameworkFromGoPackage(&p)
  }

  utils.PrintCyanInfoMsg("  Preparing project files...\n")
  err := writeFiles(&p)
  if err != nil {
    fmt.Print(err.Error())
  }
}

func getFrameworkFromGoPackage(p *Project){
  if p.Framework == "Echo" {
    utils.PrintCyanInfoMsg("  Get Echo framework dependencies...\n")
    utils.RunCommand("go", false, "get github.com/labstack/echo/v4")
  } else if p.Framework == "Fiber" {
    utils.PrintCyanInfoMsg("  Get Fiber framework dependencies...\n")
    utils.RunCommand("go", false, "get github.com/gofiber/fiber/v2")
  } 
}


func writeFiles(p *Project) error {

  var wg sync.WaitGroup

  errChan := make(chan error, len(p.Files))

  defer close(errChan)

  for _, file := range p.Files{
    wg.Add(1)
    go writeFile(file, p, &wg, errChan)
  }

  for _, file := range p.Files{
    if err := <- errChan; err != nil {
      fmt.Printf("    Error writing file %s%s: %s", file.FileName, file.Extension, err.Error())
      os.Exit(1)
    }
  }

  wg.Wait()
  return nil
}

func writeFile(f File, p *Project, wg *sync.WaitGroup, errChan chan<- error) {
  defer wg.Done()

  fileLocation := fmt.Sprintf("%s%s%s", f.FileLocation, f.FileName, f.Extension)
  fileName := fmt.Sprintf("%s%s", f.FileName, f.Extension)

	file, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
    errChan <- &errors.ClientError{
      Msg: fmt.Sprintf("    Cannot open file: %s, you need to run this command at the root of your project.", fileName),
    }
	} else {
    errChan <- nil
  }

  utils.PrintInfoMsg(fmt.Sprintf("   - Writing file %s in location => %s", fileName, fileLocation))

	defer file.Close()

	var content string

  switch fileName {

  case "main.go":
    if p.Framework == "Fiber" {
    content = fmt.Sprintf(`
    package main

    import (
      "github.com/gofiber/fiber/v2"
      h "%s/handlers"
      r "%s/router"

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
    } else { 
      content = fmt.Sprintf(`
      package main

import (
	h %s/handlers" 
	r "%s/router"
)

func main() {

	router := r.NewRouter()
	router.ServeStatic()

	//Should be in Routes
	router.HandleFunc("/", h.IndexHandler)

	router.ListenAndServe(":3000")
}`, p.Reponame, p.Reponame)
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
    if p.Framework == ""{
    content = fmt.Sprintf(`
package routers

import (
	"net/http"
	"path/filepath"
	h "%s/handlers"
)

type Router struct {
	router *http.ServeMux
}

func NewRouter() *Router {
	return &Router{router: http.NewServeMux()}
}

func (r *Router) ServeStatic() {
	staticDir := http.Dir(filepath.Join(".", "static"))
	fileServer := http.FileServer(staticDir)
	r.router.Handle("/static/", http.StripPrefix("/static/", fileServer))
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.router.HandleFunc(pattern, handler)
}

func (r *Router) ListenAndServe(port string) {
	http.ListenAndServe(port, r.router)
}`, p.Reponame)
    }else if p.Framework == "Fiber"{
      content = `//TODO WRITE ROUTER FIBER`
    }else if p.Framework == "Echo"{
      content = `//TODO WRITE ROUTER ECHO`
    }
      break

  case "user.go":
      content = `//TODO WRITE MODEL`
      break

  case "error.go":
    content = `
      package errors

import (
	"fmt"
)

type CustomError struct {
	Message string
	Status  int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func New(message string, status int) *CustomError {
	return &CustomError{Message: message, Status: status}
}`
      break

  case "utils.go":
    content = "//TODO WRITE UTILS"
      break

  case "auth-middleware.go":
    content = "//TODO WRITE MIDDLEWARES;"
      break

  case "header.templ":
    content = `
package partials

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
    content = `
package partials

templ footer(data map[string]interface{}) {
<footer>
  <p>© 2017 Company, Inc.</p>
</footer>
}`
    break

  case "base.templ":
    content = fmt.Sprintf(`
package layouts

import (
met "%s/views/metadatas"
par "%s/views/partials"
t "github.com/a-h/templ"
)

var (
Head = met.Head
Header = par.Header
Footer = par.Footer
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
    content = fmt.Sprintf(`
package pages

templ Index() {
<main>
  <h1 class="text-blue-400">Welcome</h1>
  <p class="text-blue-400">This is a simple example of a Go web app</p>
</main>
}

templ About(data map[string]interface{}) {
<main>
  <h1 class="text-blue-400">About</h1>
  <p class="text-blue-400">This is a simple example of a Go web app</p>
</main>
}

templ Contact(data map[string]interface{}) {
<main>
  <h1 class="text-blue-400">Contact</h1>
  <p class="text-blue-400">This is a simple example of a Go web app</p>
</main>
}`)
    break
   }

  _, err = file.WriteString(content)
  if err != nil {
      errChan <- &errors.ClientError{
        Msg: fmt.Sprintf("Cannot open file: %s, you need to run this command at the root of your project.", fileName),
      }
    } else {
    errChan <- nil
  }
}

