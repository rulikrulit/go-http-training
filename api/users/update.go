package users

import (
	"encoding/json"
	"go-http-training/api"
	"go-http-training/models"
	"go-http-training/storage"
	"log"

	"io/ioutil"
	"net/http"
)

func createUpdateHandler(s *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.URL.Query()["username"]

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed reading request params"))
			return
		}

		_, err := s.GetByUsername(string(username[0]))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed getting user for an update"))
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed reading request body"))
			return
		}

		u := models.User{}

		err = json.Unmarshal(data, &u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed parsing request body"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed retrieving user"))
			return
		}

		log.Printf("searching user %s", username)
		err = s.DeleteByUsername(string(username[0]))
		s.Set(u)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed retrieving user"))
			return
		}

		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(username[0] + " successfully updated"))
			return
		}
	}
}

// CreateUpdateRoute defines route for querying a user and then updating it
func CreateUpdateRoute(s *storage.Storage) api.Route {
	return api.Route{
		Path:    "/update",
		Method:  "PUT",
		Handler: createUpdateHandler(s),
	}
}
