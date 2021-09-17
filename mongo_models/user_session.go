package mongo_models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserSession struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Session   string        `json:"session"`
	UserId    string        `json:"user_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
}
