package httpd

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	server *mux.Router
}

func NewRouter() *Router {
	return &Router{server: mux.NewRouter()}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.server.ServeHTTP(w, req)
}

func (r *Router) Handle(path string, handlerFunc http.HandlerFunc) *Router {
	r.server.HandleFunc(path, handlerFunc)
	return r
}
