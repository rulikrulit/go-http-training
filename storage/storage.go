package storage

import (
	"fmt"
	"go-http-training/models"
	"log"
)

// Setter for writing into the storage
type Setter interface {
	Set(models.User) error
}

// Getter for finding and reading in storage
// Returns the first found item
type Getter interface {
	GetByUsername(username string) (models.User, error)
}

// Deleter for finding and reading in storage
// Deletes the first found item
type Deleter interface {
	DeleteByUsername(username string) error
}

// Storage main type combining all storage methods
type Storage map[string]models.User

// Set method implementation
func (s Storage) Set(user models.User) error {
	if _, ok := s[user.Username]; !ok {
		log.Println("storing user", user)
		s[user.Username] = user
		return nil
	}
	return fmt.Errorf("User already exists")
}

// GetByUsername method implementation
func (s Storage) GetByUsername(username string) (models.User, error) {
	var user models.User
	if user, ok := s[username]; ok {
		return user, nil
	}
	return user, fmt.Errorf("User not found")
}

// DeleteByUsername method implementation
func (s Storage) DeleteByUsername(username string) error {
	if _, ok := s[username]; ok {
		delete(s, username)
		return nil
	}
	return fmt.Errorf("User not found")
}

// CreateStorage sets the new storage
func CreateStorage() Storage {
	return Storage{}
}
