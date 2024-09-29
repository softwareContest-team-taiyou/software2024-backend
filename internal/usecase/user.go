package usecase

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	CheckUserExists(ctx context.Context, id string) (bool, error)
}

type UserUsecase struct {
	ur UserRepository
}

func NewUserUsecase(userRepository UserRepository) *UserUsecase {
	return &UserUsecase{ur: userRepository}
}

func(uu *UserUsecase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	user, err := uu.ur.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func(uu *UserUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	// もしuserがあればtrueを返す
	isUser, err := uu.ur.CheckUserExists(ctx, user.ID)
	if err != nil {
		return err
	}
	// userがいればtrueを返す
	if isUser {
		return nil
	}
	// なければfalseを返す
	if err := uu.ur.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}