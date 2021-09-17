package models

import "time"

type UserSession struct {
	ID         int       `json:"id"`
	UserId     int       `json:"user_id"`
	ClientType string    `json:"client_type"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
