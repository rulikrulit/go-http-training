package users

import (
	"go-http-training/api"
	"go-http-training/storage"
	"log"

	"encoding/json"
	"net/http"
)

func createGetHandler(s *storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.URL.Query()["username"]

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed reading request params"))
			return
		}

		log.Printf("searching user %s", username)
		user, err := s.Get("username", string(username[0]))
		log.Println("user found", user, err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed retrieving user"))
			return
		}

		resp, err := json.Marshal(user)

		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(resp))
		}
	}
}

// CreateGetRoute defines route for querying a user
func CreateGetRoute(s *storage.Storage) api.Route {
	return api.Route{
		Path:    "/get",
		Method:  "GET",
		Handler: createGetHandler(s),
	}
}
