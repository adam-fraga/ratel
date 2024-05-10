package handlers

import (
	"fmt"
)

func CreateGenericMiddleware() {
	fmt.Println("Creating a new middleware")
}

func CreateAuthMiddleware(name string) {
	fmt.Println("Creating a new auth middleware with name: ", name)
}
