package repository

import (
	"context"
	"time"

	"github.com/readtruyen/go-novelstory-api/domain"
	"github.com/readtruyen/go-novelstory-api/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	currentTime := primitive.NewDateTimeFromTime(time.Now())

	// Set default fields
	if user.Avatar == "" {
		user.Avatar = "https://img.freepik.com/premium-vector/man-avatar-profile-round-icon_24640-14044.jpg"
	}

	user.Active = true
	user.Common.CreatedAt = currentTime
	user.Common.UpdatedAt = currentTime

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Update(c context.Context, id string, profile *domain.UpdateProfile) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	profile.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err = collection.UpdateOne(c, bson.M{"_id": idHex}, bson.M{
		"$set": profile,
	})

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
