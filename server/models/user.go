package models

import "time"

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Favorites []string  `bson:"favorites"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Name     string     `bson:"name" json:"name"`
}
