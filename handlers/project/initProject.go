/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package project contains handlers to execute the logic of the project system of ratel web framework
*/

package project

import "fmt"

// Check first if the project is created (Presence of the project files)
// Ask user wich framework he want to use between None, Echo or Fiber

// Init the project creation process
func InitProject(reponame string, framework string) {
  fmt.Println("init project with repo: ", reponame)
  fmt.Println("Framework: ", framework)
}

