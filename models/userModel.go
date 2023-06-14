package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"id"`
	First_name    *string            `json:"first_name" validate:"required",min=2,max=100`
	Last_name     *string
	Password      *string
	Email         *string
	Phone         *string
	Token         *string
	User_type     *string
	Refresh_token *string
	Created_at    time.Time
	Updated_at    time.Time
	User_id       string
}
