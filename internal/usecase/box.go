package usecase

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type BoxUsecase struct {
	br BoxRepository
}

func NewBoxUsecase(br BoxRepository) *BoxUsecase {
	return &BoxUsecase{
		br: br,
	}
}

func (bu *BoxUsecase) IsLock(ctx context.Context, UserId string) (*domain.Box, error) {
	return bu.br.IsLock(ctx, UserId)
}

func (bu *BoxUsecase) Lock(ctx context.Context, UserId string) error {
	return bu.br.Lock(ctx, UserId)
}

func (bu *BoxUsecase) Unlock(ctx context.Context, UserId string) error {
	return bu.br.Unlock(ctx, UserId)
}