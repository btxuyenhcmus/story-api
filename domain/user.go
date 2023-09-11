package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Avatar      string             `bson:"avatar"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	DateOfBirth string             `bson:"date_of_birth"`
	Gender      string             `bson:"gender"`
	PhoneNumber string             `bson:"phone_number"`
	Address     Address            `bson:"inline"`
	Active      bool               `bson:"active"`
	Roles       []string           `bson:"roles"`
	Common      Common             `bson:"inline"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Update(c context.Context, id string, profile *UpdateProfile) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
