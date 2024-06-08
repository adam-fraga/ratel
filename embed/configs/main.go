package main

import (
	h "github.com/your-name/your-appname/handlers" // Change path to match your project path and uncomment
	r "github.com/your-name/your-appname/router"   // Change path to match your project path and uncomment
)

func main() {

	router := r.NewRouter()
	router.ServeStatic()

	//Should be in Routes
	router.HandleFunc("/", h.IndexHandler)

	router.ListenAndServe(":3000")
}
