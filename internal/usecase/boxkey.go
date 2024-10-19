package usecase

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type KeyRepository interface {
	CreateKey(ctx context.Context, key *domain.Key,userId string) error
}
type BoxRepository interface {
	CreateBox(ctx context.Context, box *domain.Box,userId string) error
	IsLock(ctx context.Context, UserId string) (*domain.Box, error)
	Lock (ctx context.Context, UserId string) error
	Unlock (ctx context.Context, UserId string) error
}

type BoxKeyRepository interface {
	CreateBoxKey(ctx context.Context, box *domain.Box, key *domain.Key) error
}

type BoxKeyUseCase struct {
	ur UserRepository
	bkr BoxKeyRepository
	kr KeyRepository
	br BoxRepository
}

func NewBoxKeyUsecase(userRepository UserRepository, boxKeyRepository BoxKeyRepository, keyRepository KeyRepository,boxRepository BoxRepository ) *BoxKeyUseCase {
	return &BoxKeyUseCase{ur: userRepository, bkr: boxKeyRepository, kr: keyRepository, br: boxRepository}
}


func (bu *BoxKeyUseCase) InitCreateBoxKey(ctx context.Context, box *domain.Box, key *domain.Key, userID string) error {
	
	if err := bu.br.CreateBox(ctx, box, userID); err != nil {
		return err
	}
	if err := bu.kr.CreateKey(ctx, key,userID); err != nil {
		return err
	}
	if err := bu.bkr.CreateBoxKey(ctx, box, key); err != nil {
		return err
	}
    user, err := bu.ur.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	UpdateUser := &domain.User{
		ID: user.ID,
		Name: user.Name,
		IsInit: true,
	}
	if err := bu.ur.UpdateUser(ctx, UpdateUser); err != nil {
		return err
	}
	return nil
}

