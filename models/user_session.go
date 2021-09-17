package models

import "time"

type UserSession struct {
	ID         int       `json:"id"`
	UserId     int       `json:"userId"`
	ClientType string    `json:"clientType"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
