package storage

import (
	"fmt"
	"log"
	"reflect"
)

// Item custom type Data in storage
type Item interface{}

// FieldValue custom field value in storage
type FieldValue interface{}

// Setter for writing into the storage
type Setter interface {
	Set(Item) error
}

// Getter for finding and reading in storage
// Returns the first found item
type Getter interface {
	Get(field string, value FieldValue) (Item, error)
}

// Deleter for finding and reading in storage
// Deletes the first found item
type Deleter interface {
	Delete(field string, value FieldValue) error
}

// Storage main type combining all storage methods
type Storage struct {
	Data []Item
	Setter
	Getter
	Deleter
}

// Set method implementation
func (s *Storage) Set(i Item) error {
	log.Println(i)
	s.Data = append(s.Data, i)
	return nil
}

// Get method implementation
func (s Storage) Get(f string, v FieldValue) (Item, error) {
	for _, i := range s.Data {
		r := reflect.ValueOf(i)

		// Magic! Getting first field and casting string type onto it
		if reflect.Indirect(r).FieldByIndex([]int{0}).Interface().(string) == v {
			return i, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}

// Delete method implementation
func (s *Storage) Delete(f string, v FieldValue) error {
	for i, item := range s.Data {
		r := reflect.ValueOf(item)

		if reflect.Indirect(r).FieldByIndex([]int{0}).Interface().(string) == v {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

// CreateStorage sets the new storage
func CreateStorage() Storage {
	return Storage{}
}
