package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Avatar      string   `json:"avatar"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	DateOfBirth string   `json:"date_of_birth"`
	Gender      string   `json:"gender"`
	PhoneNumber string   `json:"phone_number"`
	Roles       []string `json:"roles"`
}

type UpdateProfile struct {
	Avatar      string             `bson:"avatar" form:"avatar"`
	Name        string             `bson:"name" form:"name"`
	DateOfBirth string             `bson:"date_of_birth" form:"date_of_birth"`
	Gender      string             `bson:"gender" form:"gender"`
	PhoneNumber string             `bson:"phone_number" form:"phone_number"`
	UpdatedAt   primitive.DateTime `bson:"updated_at"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
	UpdateProfile(c context.Context, userID string, profile *UpdateProfile) error
}
