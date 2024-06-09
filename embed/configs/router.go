package routers

import (
	h "github.com/your-name/app-name/handlers"
	"net/http"
	"path/filepath"
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
}
