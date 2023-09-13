package repository

import (
	"context"
	"errors"
	"time"

	"github.com/readtruyen/go-novelstory-api/domain"
	"github.com/readtruyen/go-novelstory-api/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type chapterRepository struct {
	database   mongo.Database
	collection string
}

func NewChapterRepository(db mongo.Database, collection string) domain.ChapterRepository {
	return &chapterRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *chapterRepository) Create(c context.Context, chapter *domain.Chapter) error {
	collection := cr.database.Collection(cr.collection)

	var existChapter domain.Chapter
	err := collection.FindOne(c, bson.M{"id": chapter.SourceID}).Decode(&existChapter)
	if err == nil {
		return errors.New("CHAPTER EXISTED")
	}

	// Setup default fields
	currentTime := primitive.NewDateTimeFromTime(time.Now())
	chapter.Common.CreatedAt = currentTime
	chapter.Common.UpdatedAt = currentTime

	_, err = collection.InsertOne(c, chapter)

	return err
}

func (cr *chapterRepository) Update(c context.Context, chapterID string, updateChapter *domain.UpdateChapter) error {
	return nil
}

func (cr *chapterRepository) GetChaptersByStoryId(c context.Context, storyID string) ([]domain.ShortResponseChapter, error) {
	collection := cr.database.Collection(cr.collection)

	var chapters []domain.ShortResponseChapter
	storyObjectId, err := primitive.ObjectIDFromHex(storyID)
	if err != nil {
		return []domain.ShortResponseChapter{}, err
	}

	cursor, err := collection.Find(c, bson.M{"story_id": storyObjectId})

	if err != nil {
		return []domain.ShortResponseChapter{}, err
	}

	err = cursor.All(c, &chapters)
	if chapters == nil {
		return []domain.ShortResponseChapter{}, err
	}
	return chapters, err
}

func (cr *chapterRepository) DownloadChapterByStoryId(c context.Context, storyID string) ([]domain.Chapter, error) {
	collection := cr.database.Collection(cr.collection)

	var chapters []domain.Chapter
	storyObjectId, err := primitive.ObjectIDFromHex(storyID)
	if err != nil {
		return []domain.Chapter{}, err
	}

	cursor, err := collection.Find(c, bson.M{"story_id": storyObjectId})

	if err != nil {
		return []domain.Chapter{}, err
	}

	err = cursor.All(c, &chapters)
	if chapters == nil {
		return []domain.Chapter{}, err
	}
	return chapters, err
}

func (cr *chapterRepository) GetChaptersByStoryIdPagination(c context.Context, storyID string, page int, perPage int) ([]domain.Chapter, domain.Pagination, error) {
	return []domain.Chapter{}, domain.Pagination{}, nil
}

func (cr *chapterRepository) GetChapterById(c context.Context, chapterID string) (domain.Chapter, error) {
	collection := cr.database.Collection(cr.collection)

	var chapter domain.Chapter

	idHex, err := primitive.ObjectIDFromHex(chapterID)
	if err != nil {
		return chapter, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&chapter)
	return chapter, err
}

func (cr *chapterRepository) DeleteChapterById(c context.Context, chapterID string) error {
	return nil
}
