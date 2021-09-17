package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserId struct {
	ID int `json:"id"`
}
