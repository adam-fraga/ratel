/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package project contains handlers to execute the logic of the project system of ratel web framework
*/

package project

import (
	"fmt"

	"github.com/adam-fraga/ratel/utils"
)

// Check first if the project is created (Presence of the project files)
// Ask user wich framework he want to use between None, Echo or Fiber

// Init the project creation process
func InitProject(reponame string, framework string) {
  if framework == "" {
    utils.RunCommandWithOutput("go", fmt.Sprintf("mod init %s", reponame))
    utils.PrintSuccessMsg(fmt.Sprintf("Successfully initialized repo: %s !", reponame))
    return
  }

  if framework == "echo" || framework == "fiber"{
    utils.RunCommandWithOutput("go", fmt.Sprintf("mod init %s", reponame))
    utils.PrintSuccessMsg(fmt.Sprintf("Successfully initialized repo: %s with %s framework !", reponame, framework))
    getFrameworkFromGoPackage(framework)
  }
}

func getFrameworkFromGoPackage(framework string){
  if framework == "echo" {
    utils.PrintInfoMsg("Get echo framework from go package...")
    utils.RunCommandWithOutput("go", "get github.com/labstack/echo/v4")
    return
  }
  if framework == "fiber" {
    utils.PrintInfoMsg("Get echo framework from go package...")
    utils.RunCommandWithOutput("go", "get github.com/gofiber/fiber/v2")
  }

}
