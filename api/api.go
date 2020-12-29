package api

import (
	"log"
	"net/http"
)

// RoutesCreator is the server interface
type RoutesCreator interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

// Route defining structure of a route with a handler
type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// CreateRoute hepler to create route onto the predefined server
func CreateRoute(srv RoutesCreator, rt Route, ctx string) {
	path := ctx + rt.Path
	log.Println("created path " + rt.Method + " " + path)
	srv.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		log.Println("called path " + path)
		if r.Method == rt.Method {
			rt.Handler(w, r)
		} else {
			log.Println("No handling of multiple methods for " + path + ". Use " + rt.Method)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		}
	})
}
