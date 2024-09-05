/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package project contains handlers to execute the logic of the project system of ratel web framework
*/

package project

import (
	"fmt"

  s "strings"
	"github.com/adam-fraga/ratel/utils"
)

// Check first if the project is created (Presence of the project files)
// Ask user wich framework he want to use between None, Echo or Fiber

// Init the project creation process
func InitProject(reponame string, framework string) {
  p := Project{}

  if reponame == "" || !s.HasPrefix("github/", reponame){
    utils.PrintErrorMsg("Repo name is not well formatted: \"github/your-name/repo\"")
    return
  }

  p.Reponame = reponame

  if framework == "" {
    utils.RunCommandWithOutput("go", fmt.Sprintf("mod init %s", reponame))
    utils.PrintSuccessMsg(fmt.Sprintf("Successfully initialized repo: %s !", reponame))
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
    return
  } else if p.Framework == "Fiber" {
    utils.PrintInfoMsg("Get echo framework from go package...")
    utils.RunCommandWithOutput("go", "get github.com/gofiber/fiber/v2")
    return
  }
}

func rewriteGoMainFileAccordingToFramework(p *Project)  {

}
  
