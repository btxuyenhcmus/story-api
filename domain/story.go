package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionStory = "stories"
)

type Story struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	SourceID      int                `bson:"id" json:"-" form:"source_id"`
	Title         string             `bson:"title" binding:"required" json:"title" form:"title"`
	Image         string             `bson:"image" binding:"required" json:"image" form:"image"`
	IsFull        bool               `bson:"is_full" binding:"required" json:"is_full" form:"is_full"`
	Link          string             `bson:"link" binding:"required" json:"link" form:"link"`
	Status        string             `bson:"status" binding:"required" json:"status" form:"status"`
	TotalChapters int                `bson:"total_chapters" json:"total_chapters" form:"total_chapters"`
	Authors       []string           `bson:"authors" json:"authors"`
	Categories    []string           `bson:"categories" json:"categories"`
	Description   string             `bson:"description" json:"description" form:"description"`
	Common        Common             `bson:"inline"`
}

type UpdateStory struct {
	Title       string   `bson:"title" form:"title"`
	Image       string   `bson:"image" form:"image"`
	IsFull      bool     `bson:"is_full" form:"is_full"`
	Link        string   `bson:"link" form:"link"`
	Status      string   `bson:"status" form:"status"`
	Authors     []string `bson:"authors" form:"authors"`
	Categories  []string `bson:"categories" form:"categories"`
	Description string   `bson:"description" form:"description"`
}

type ShortResponseStory struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Title  string             `bson:"title" json:"title"`
	Image  string             `bson:"image" json:"image"`
	IsFull bool               `bson:"is_full" json:"is_full" `
}

type StoryRepository interface {
	Create(c context.Context, story *Story) error
	Update(c context.Context, storyID string, updateStory *UpdateStory) error
	GetStories(c context.Context) ([]ShortResponseStory, error)
	GetStoriesPagination(c context.Context, page int, perPage int) ([]ShortResponseStory, Pagination, error)
	GetStoryById(c context.Context, storyID string) (Story, error)
	DeleteStoryByID(c context.Context, storyID string) error
}

type StoryUseCase interface {
	Create(c context.Context, story *Story) error
	GetStories(c context.Context) ([]ShortResponseStory, error)
	GetStoriesPagination(c context.Context, page int, perPage int) ([]ShortResponseStory, Pagination, error)
	GetStoryById(c context.Context, storyID string) (Story, error)
}
