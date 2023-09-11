package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Common struct {
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type Address struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	State   string `bson:"state"`
	ZipCode string `bson:"zip_code"`
	Country string `bson:"country"`
}
