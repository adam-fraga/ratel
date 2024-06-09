package handlers

import (
	// t "github.com/a-h/templ"
	// pages "github.com/adam-fraga/myapp/views/pages" #Change path to match your project and uncomment
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	homeData := make(map[string]interface{})
	homeData["name"] = "John"
	// homePage := pages.Index(homeData)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	// t.Handler(homePage).ServeHTTP(w, r)
}
