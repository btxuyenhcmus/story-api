package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionChapter = "chapters"
)

type Chapter struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	SourceID      int                `bson:"id" json:"source_id" form:"source_id"`
	StoryID       primitive.ObjectID `bson:"story_id" json:"story_id"`
	ChapterNext   primitive.ObjectID `bson:"chapter_next,omitempty" json:"chapter_next,omitempty" form:"chapter_next"`
	ChapterNextID int                `bson:"chapter_next_id" json:"-" form:"chapter_next_id"`
	ChapterPrev   primitive.ObjectID `bson:"chapter_prev,omitempty" json:"chapter_prev,omitempty" form:"chapter_prev"`
	ChapterPrevID int                `bson:"chapter_prev_id" json:"-" form:"chapter_prev_id"`
	Title         string             `bson:"title" binding:"required" json:"title" form:"title"`
	Content       string             `bson:"content" json:"content" form:"content"`
	Number        int                `bson:"number" json:"number" form:"number"`
	Common        Common             `bson:"inline"`
}

type CreateChapter struct {
	SourceID      int    `bson:"id" json:"source_id" form:"source_id"`
	StoryID       string `bson:"story_id" json:"story_id"`
	ChapterNextID int    `bson:"chapter_next_id" json:"-" form:"chapter_next_id"`
	ChapterPrev   string `bson:"chapter_prev,omitempty" json:"chapter_prev,omitempty" form:"chapter_prev"`
	ChapterPrevID int    `bson:"chapter_prev_id" json:"-" form:"chapter_prev_id"`
	Title         string `bson:"title" binding:"required" json:"title" form:"title"`
	Content       string `bson:"content" json:"content" form:"content"`
	Number        int    `bson:"number" json:"number" form:"number"`
}

type UpdateChapter struct {
	Title   string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
}

type ShortResponseChapter struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" binding:"required" json:"title" form:"title"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type ChapterRepository interface {
	Create(c context.Context, chapter *Chapter) error
	Update(c context.Context, chapterID string, updateChapter *UpdateChapter) error
	GetChaptersByStoryId(c context.Context, storyID string) ([]ShortResponseChapter, error)
	DownloadChapterByStoryId(c context.Context, storyID string) ([]Chapter, error)
	GetChaptersByStoryIdPagination(c context.Context, storyID string, page int, perPage int) ([]Chapter, Pagination, error)
	GetChapterById(c context.Context, chapterID string) (Chapter, error)
	DeleteChapterById(c context.Context, chapterID string) error
}

type ChapterUseCase interface {
	Create(c context.Context, chapter *Chapter) error
	GetChaptersByStoryId(c context.Context, storyID string) ([]ShortResponseChapter, error)
	DownloadChapterByStoryId(c context.Context, storyID string) ([]Chapter, error)
	GetChapterById(c context.Context, chapterID string) (Chapter, error)
}
