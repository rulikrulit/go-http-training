package users

import (
	"go-http-training/api"
	"net/http"
)

func createHanlder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// CreateRoute defines route for users lists
var CreateRoute = api.Route{
	Path:    "/create",
	Method:  "POST",
	Handler: createHanlder,
}
