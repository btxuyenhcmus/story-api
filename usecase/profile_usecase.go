package usecase

import (
	"context"
	"time"

	"readtruyen-api/domain"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Avatar: user.Avatar, Name: user.Name, Email: user.Email, DateOfBirth: user.DateOfBirth, Gender: user.Gender, PhoneNumber: user.PhoneNumber, Roles: user.Roles}, nil
}

func (pu *profileUsecase) UpdateProfile(c context.Context, userID string, profile *domain.UpdateProfile) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	err := pu.userRepository.Update(ctx, userID, profile)
	return err
}
