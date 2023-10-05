package usecase

import (
	"context"
	"time"

	"github.com/readtruyen/go-novelstory-api/domain"
)

type carouselUseCase struct {
	carouselRepository domain.CarouselRepository
	contextTimeout     time.Duration
}

func NewCarouselUseCase(carouselRepository domain.CarouselRepository, timeout time.Duration) domain.CarouselUseCase {
	return &carouselUseCase{
		carouselRepository: carouselRepository,
		contextTimeout:     timeout,
	}
}

func (cu *carouselUseCase) Create(c context.Context, carousel *domain.Carousel) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.carouselRepository.Create(ctx, carousel)
}

func (cu *carouselUseCase) GetCarousels(c context.Context) ([]domain.Carousel, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.carouselRepository.GetCarousels(ctx)
}

func (cu *carouselUseCase) DeleteCarouselById(c context.Context, carouselID string) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.carouselRepository.DeleteCarouselById(ctx, carouselID)
}
