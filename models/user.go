package models

type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
}
