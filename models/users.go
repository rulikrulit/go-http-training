package models

// User describes base user model
type User struct {
	Username string `json:"username"`
	Name     string `json:"full_name"`
	Age      uint   `json:"age"`
	Gender   string `json:"gender"`
}
