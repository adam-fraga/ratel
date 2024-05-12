package handlers

import (
	"fmt"
)

func CreateGenericMiddleware(name string) {
	fmt.Println("Creating a new middleware")
}

func CreateAuthMiddleware(name string) {
	fmt.Println("Creating a new auth middleware with name: ", name)
}

func ListMiddlewares() {
	fmt.Println("Listing all middlewares")
}
