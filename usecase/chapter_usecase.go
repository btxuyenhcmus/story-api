package usecase

import (
	"context"
	"time"

	"github.com/readtruyen/go-novelstory-api/domain"
)

type chapterUseCase struct {
	chapterRepository domain.ChapterRepository
	contextTimeout    time.Duration
}

func NewChapterUseCase(chapterRepository domain.ChapterRepository, timeout time.Duration) domain.ChapterUseCase {
	return &chapterUseCase{
		chapterRepository: chapterRepository,
		contextTimeout:    timeout,
	}
}

func (cu *chapterUseCase) Create(c context.Context, chapter *domain.Chapter) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chapterRepository.Create(ctx, chapter)
}

func (cu *chapterUseCase) GetChaptersByStoryId(c context.Context, storyID string) ([]domain.ShortResponseChapter, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chapterRepository.GetChaptersByStoryId(ctx, storyID)
}

func (cu *chapterUseCase) DownloadChapterByStoryId(c context.Context, storyID string) ([]domain.Chapter, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chapterRepository.DownloadChapterByStoryId(ctx, storyID)
}

func (cu *chapterUseCase) GetChapterById(c context.Context, chapterID string) (domain.Chapter, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chapterRepository.GetChapterById(ctx, chapterID)
}
