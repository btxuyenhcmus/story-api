package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCarousel = "carousels"
)

type Carousel struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Image           string             `bson:"image" binding:"required" json:"image" form:"image"`
	DestinationId   int                `bson:"destination_id" binding:"required" json:"destination_id" form:"destination_id"`
	DestinationType string             `bson:"destination_type" json:"destination_type" form:"destination_type"`
}

type CarouselRepository interface {
	Create(c context.Context, carousel *Carousel) error
	GetCarousels(c context.Context) ([]Carousel, error)
	DeleteCarouselById(c context.Context, carouselID string) error
}

type CarouselUseCase interface {
	Create(c context.Context, carousel *Carousel) error
	GetCarousels(c context.Context) ([]Carousel, error)
	DeleteCarouselById(c context.Context, carouselID string) error
}
