package users

import (
	"go-http-training/api"
	"go-http-training/models"
	"go-http-training/storage"
	"io/ioutil"

	"encoding/json"
	"net/http"
)

func createAddHandler(s *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed reading request body"))
			return
		}
		user := models.User{}
		err = json.Unmarshal(data, &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed parsing request body"))
			return
		}

		s.Set(user)

		resp, err := json.Marshal(user)

		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(resp))
		}
	}
}

// CreateAddRoute defines route for creating a user
func CreateAddRoute(s *storage.Storage) api.Route {
	return api.Route{
		Path:    "/create",
		Method:  "POST",
		Handler: createAddHandler(s),
	}
}
