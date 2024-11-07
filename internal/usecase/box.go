package usecase

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)
type SlackService interface {
	SendMessage(ctx context.Context, message string) error
}
type BoxUsecase struct {
	br BoxRepository
	ss SlackService
}


func NewBoxUsecase(br BoxRepository,ss SlackService) *BoxUsecase {
	return &BoxUsecase{
		br: br,
		ss: ss,
	}
}

func (bu *BoxUsecase) IsLock(ctx context.Context, UserId string) (*domain.Box, error) {
	return bu.br.IsLock(ctx, UserId)
}

func (bu *BoxUsecase) Lock(ctx context.Context, UserId string) error {
	box,err := bu.br.IsLock(ctx, UserId)
	if err != nil {
		return err
	}
	if box.IsLock {
		return nil
		}

	if err := bu.ss.SendMessage(ctx, "lock"+box.ID); err != nil {
		return err
	}
	return bu.br.Lock(ctx, UserId)
}

func (bu *BoxUsecase) Unlock(ctx context.Context, UserId string) error {
	box,err := bu.br.IsLock(ctx, UserId)
	if err != nil {
		return err
	}
	if !box.IsLock {
		return nil
	}
	if err := bu.ss.SendMessage(ctx, "unlock"+box.ID); err != nil {
		return err
	}
	return bu.br.Unlock(ctx, UserId)
}