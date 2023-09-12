package repository

import (
	"context"
	"time"

	"github.com/readtruyen/go-novelstory-api/domain"
	"github.com/readtruyen/go-novelstory-api/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type storyRepository struct {
	database   mongo.Database
	collection string
}

func NewStoryRepository(db mongo.Database, collection string) domain.StoryRepository {
	return &storyRepository{
		database:   db,
		collection: collection,
	}
}

func (sr *storyRepository) Create(c context.Context, story *domain.Story) error {
	collection := sr.database.Collection(sr.collection)

	// Setup default fields
	currentTime := primitive.NewDateTimeFromTime(time.Now())
	story.Common.CreatedAt = currentTime
	story.Common.UpdatedAt = currentTime

	_, err := collection.InsertOne(c, story)

	return err
}

func (sr *storyRepository) Update(c context.Context, storyID string, updateStory *domain.UpdateStory) error {
	return nil
}

func (sr *storyRepository) GetStories(c context.Context) ([]domain.ShortResponseStory, error) {
	collection := sr.database.Collection(sr.collection)

	var stories []domain.ShortResponseStory

	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &stories)
	if stories == nil {
		return []domain.ShortResponseStory{}, err
	}

	return stories, err
}

func (sr *storyRepository) GetStoriesPagination(c context.Context, page int, perPage int) ([]domain.ShortResponseStory, domain.Pagination, error) {
	collection := sr.database.Collection(sr.collection)

	var stories []domain.ShortResponseStory

	skip := (page - 1) * perPage

	cursor, err := collection.Aggregate(c, []bson.M{
		{"$skip": skip},
		{"$limit": perPage},
	})
	if err != nil {
		return nil, domain.Pagination{}, err
	}

	for cursor.Next(c) {
		var items domain.ShortResponseStory
		if err := cursor.Decode(&items); err != nil {
			return nil, domain.Pagination{}, err
		}
		stories = append(stories, items)
	}

	totalCount, err := collection.CountDocuments(c, bson.M{})
	if err != nil {
		return nil, domain.Pagination{}, err
	}

	totalPages := totalCount / int64(perPage)
	if totalPages%int64(perPage) > 0 {
		totalPages++
	}

	return stories, domain.Pagination{TotalPages: int(totalPages), CurrentPage: page, Total: int(totalCount), PerPage: perPage, Count: len(stories), Next: page + 1}, err
}

func (sr *storyRepository) GetStoryById(c context.Context, storyID string) (domain.Story, error) {
	collection := sr.database.Collection(sr.collection)

	var story domain.Story

	idHex, err := primitive.ObjectIDFromHex(storyID)
	if err != nil {
		return story, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&story)
	return story, err
}

func (sr *storyRepository) DeleteStoryByID(c context.Context, storyID string) error {
	return nil
}
