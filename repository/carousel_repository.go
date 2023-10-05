package repository

import (
	"context"

	"github.com/readtruyen/go-novelstory-api/domain"
	"github.com/readtruyen/go-novelstory-api/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type carouselRepository struct {
	database   mongo.Database
	collection string
}

func NewCarouselRepository(db mongo.Database, collection string) domain.CarouselRepository {
	return &carouselRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *carouselRepository) Create(c context.Context, carousel *domain.Carousel) error {
	collection := cr.database.Collection(cr.collection)

	_, err := collection.InsertOne(c, carousel)

	return err
}

func (cr *carouselRepository) GetCarousels(c context.Context) ([]domain.Carousel, error) {
	collection := cr.database.Collection(cr.collection)

	var carousels []domain.Carousel

	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &carousels)
	if carousels == nil {
		return []domain.Carousel{}, err
	}

	return carousels, err
}

func (cr *carouselRepository) DeleteCarouselById(c context.Context, carouselID string) error {
	return nil
}
