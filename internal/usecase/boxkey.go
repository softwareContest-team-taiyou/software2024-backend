package usecase

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type KeyRepository interface {
	CreateKey(ctx context.Context, key *domain.Key) error
}
type BoxRepository interface {
	CreateBox(ctx context.Context, box *domain.Box) error
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
	if err := bu.br.CreateBox(ctx, box); err != nil {
		return err
	}
	if err := bu.kr.CreateKey(ctx, key); err != nil {
		return err
	}
	if err := bu.bkr.CreateBoxKey(ctx, box, key); err != nil {
		return err
	}
	return nil
}

