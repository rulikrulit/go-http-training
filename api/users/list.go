package users

import (
	"go-http-training/api"
	"go-http-training/storage"

	"encoding/json"
	"net/http"
)

func createListHandler(s storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp, err := json.Marshal(s.Data)

		if err == nil {
			w.Write([]byte(resp))
		}
	}
}

// CreateListRoute defines route for users lists
func CreateListRoute(s storage.Storage) api.Route {
	return api.Route{
		Path:    "/list",
		Method:  "GET",
		Handler: createListHandler(s),
	}
}
