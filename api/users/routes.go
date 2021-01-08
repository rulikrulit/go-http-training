package users

import (
	"go-http-training/api"
	"go-http-training/storage"
)

const namespace = "/users"

var store = storage.CreateStorage()

// SetUserRoutes creates routes for users
func SetUserRoutes(srv api.RoutesCreator) {
	api.CreateRoute(srv, CreateListRoute(&store), namespace)
	api.CreateRoute(srv, CreateAddRoute(&store), namespace)
	api.CreateRoute(srv, CreateGetRoute(&store), namespace)
	api.CreateRoute(srv, CreateUpdateRoute(&store), namespace)
}
