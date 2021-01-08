package storage_test

import (
	"testing"

	"go-http-training/models"
	"go-http-training/storage"
)

var testUser = models.User{
	Username: "fancy_trip",
	Name:     "Ivan Susanin",
	Age:      6,
	Gender:   "male",
}

func TestAdd(t *testing.T) {
	s := storage.CreateStorage()

	if err := s.Set(testUser); err != nil {
		t.Error("couldn't create a user")
	}
}

func TestAddDuplicate(t *testing.T) {
	s := storage.CreateStorage()

	if err := s.Set(testUser); err != nil {
		t.Error("couldn't create a user")
	}

	if err := s.Set(testUser); err == nil {
		t.Error("no error found while creating a duplicate")
	}

}

func TestFind(t *testing.T) {
	s := storage.CreateStorage()

	if err := s.Set(testUser); err != nil {
		t.Error("couldn't create a user")
	}

	if user, err := s.GetByUsername(testUser.Username); err != nil {
		t.Error("failed finding user")
	} else {
		if user.Username != testUser.Username {
			t.Error("failed finding correct user")
		}
	}
}

func TestDelete(t *testing.T) {
	s := storage.CreateStorage()

	if err := s.Set(testUser); err != nil {
		t.Error("couldn't create a user")
	}

	if _, err := s.GetByUsername(testUser.Username); err != nil {
		t.Error("failed finding user")
	}

	if err := s.DeleteByUsername(testUser.Username); err != nil {
		t.Error("failed deleting user")
	}

	if _, err := s.GetByUsername(testUser.Username); err == nil {
		t.Error("Failed deleting user. User still exists in storage")
	}

}
