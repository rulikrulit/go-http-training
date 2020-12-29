package users

import (
	"go-http-training/api"
	"go-http-training/storage"

	"encoding/json"
	"net/http"
)

func createListHandler(s *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := json.Marshal(s.Data)

		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(resp))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed response formatting"))
		}
	}
}

// CreateListRoute defines route for users lists
func CreateListRoute(s *storage.Storage) api.Route {
	return api.Route{
		Path:    "/list",
		Method:  "GET",
		Handler: createListHandler(s),
	}
}
