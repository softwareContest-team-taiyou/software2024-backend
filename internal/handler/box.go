package handler

import (
	"context"
	"errors"

	boxv1 "github.com/softwareContest-team-taiyou/software2024-backend/gen/go/v1/box"
	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type BoxUseCase interface {
	IsLock(ctx context.Context,UserId string)(*domain.Box,error) 
	Lock(ctx context.Context,UserId string) error
	Unlock(ctx context.Context,UserId string) error
}

type BoxHandler struct {
	boxv1.UnimplementedBoxServiceServer
	bu BoxUseCase
}

func NewBoxHandler(boxUseCase BoxUseCase) *BoxHandler {
	return &BoxHandler{bu: boxUseCase}
}

func (bh *BoxHandler) IsLock(ctx context.Context, req *boxv1.IsLockRequest) (*boxv1.IsLockResponse, error) {
	userId, ok := ctx.Value("uid").(string)
	if ok != true {
		return nil, errors.New("uid not found")
	}
	box,err := bh.bu.IsLock(ctx,userId)
	if err != nil {
		return nil,err
	}
	return &boxv1.IsLockResponse{
		IsLock: box.IsLock,
		Name: box.Name,
	}, nil
}

func (bh *BoxHandler) Lock(ctx context.Context, req *boxv1.LockRequest) (*boxv1.LockResponse, error) {
	userId, ok := ctx.Value("uid").(string)
	if ok != true {
		return nil, errors.New("uid not found")
	}
	if err := bh.bu.Lock(ctx,userId); err != nil {
		return nil,err
	}
	return &boxv1.LockResponse{
		Status: boxv1.Status_SUCCESS,
	}, nil
}

func (bh *BoxHandler) Unlock(ctx context.Context, req *boxv1.UnlockRequest) (*boxv1.UnlockResponse, error) {
	userId, ok := ctx.Value("uid").(string)
	if ok != true {
		return nil, errors.New("uid not found")
	}
	if err := bh.bu.Unlock(ctx,userId); err != nil {
		return nil,err
	}
	return &boxv1.UnlockResponse{
		Status: boxv1.Status_SUCCESS,
	}, nil
}



