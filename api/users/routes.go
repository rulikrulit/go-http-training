package users

import (
	"go-http-training/api"
	"go-http-training/models"
	"go-http-training/storage"
)

const namespace = "/users"

var defaultUsers = []storage.Item{
	models.User{
		Username: "1",
		Name:     "2",
		Age:      5,
		Gender:   "male",
	},
}

var store = storage.Storage{
	Data: defaultUsers,
}

// SetUserRoutes creates routes for users
func SetUserRoutes(srv api.RoutesCreator) {
	api.CreateRoute(srv, CreateListRoute(store), namespace)
	api.CreateRoute(srv, CreateRoute, namespace)
}
