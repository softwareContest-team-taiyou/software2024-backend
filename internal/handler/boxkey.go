package handler

import (
	"context"
	"errors"

	boxkeyv1 "github.com/softwareContest-team-taiyou/software2024-backend/gen/go/v1/boxkey"
	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type BoxKeyUseCase interface {
	InitCreateBoxKey(ctx context.Context, box *domain.Box, key *domain.Key,userID string) error
}

type BoxKeyHandler struct {
	boxkeyv1.UnimplementedBoxKeyServiceServer
	bu BoxKeyUseCase
}

func NewBoxKeyHandler(boxKeyUseCase BoxKeyUseCase) *BoxKeyHandler {
	return &BoxKeyHandler{bu: boxKeyUseCase}
}

func (bh *BoxKeyHandler) InitCreateBoxKey(ctx context.Context, req *boxkeyv1.InitCreateRequest) (*boxkeyv1.InitCreateResponse, error) {
	userID, ok := ctx.Value("uid").(string)
	if ok != true {
		return nil, errors.New("uid not found")
	}
	box := &domain.Box{
		ID:   req.BoxId,
		Name: req.BoxName,
		IsLock: false,
	}
	key := &domain.Key{
		ID:   req.KeyId,
		Name: req.KeyName,
	}
	if err := bh.bu.InitCreateBoxKey(ctx, box,key,userID); err != nil {
		return nil, err
	}
	return &boxkeyv1.InitCreateResponse{
		Status: boxkeyv1.Status_SUCCESS,
	}, nil
}