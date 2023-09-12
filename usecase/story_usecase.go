package usecase

import (
	"context"
	"time"

	"github.com/readtruyen/go-novelstory-api/domain"
)

type storyUseCase struct {
	storyRepository domain.StoryRepository
	contextTimeout  time.Duration
}

func NewStoryUseCase(storyRepository domain.StoryRepository, timeout time.Duration) domain.StoryUseCase {
	return &storyUseCase{
		storyRepository: storyRepository,
		contextTimeout:  timeout,
	}
}

func (su *storyUseCase) Create(c context.Context, story *domain.Story) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.storyRepository.Create(ctx, story)
}

func (su *storyUseCase) GetStories(c context.Context) ([]domain.ShortResponseStory, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.storyRepository.GetStories(ctx)
}

func (su *storyUseCase) GetStoryById(c context.Context, storyID string) (domain.Story, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.storyRepository.GetStoryById(ctx, storyID)
}
