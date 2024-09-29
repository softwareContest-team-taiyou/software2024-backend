package handler

import (
	"context"
	"errors"

	userv1 "github.com/softwareContest-team-taiyou/software2024-backend/gen/go/v1/user"
	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)


type UserUsecase interface {
	GetUser(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
}
type UserHandler struct {
	userv1.UnimplementedUserServiceServer
	uu UserUsecase
}

func NewUserHandler(userUsecase UserUsecase) *UserHandler {
	return &UserHandler{uu: userUsecase}
}

func (uh *UserHandler) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	// コンテキストからユーザIDを取得
	userID, ok := ctx.Value("uid").(string)
	if ok != true {
		return nil, errors.New("uid not found")
	}
	// usecaseを呼び出し
	user, err := uh.uu.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &userv1.GetUserResponse{
		Name : user.Name,
	}, nil
}

func (uh *UserHandler) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	userID, ok := ctx.Value("uid").(string)
	if ok != true {
		return nil, errors.New("uid not found")
	}
	user := &domain.User{
		ID:   userID,
		Name: "",
	}
	if err := uh.uu.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &userv1.CreateUserResponse{
		Status: userv1.Status_SUCCESS,
	}, nil
}

func (uh *UserHandler) UpdateUser(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UpdateUserResponse, error) {
	userID, ok := ctx.Value("uid").(string)
	if ok != true {
		return nil, errors.New("uid not found")
	}
	user := &domain.User{
		ID:   userID,
		Name: req.Name,
	}
	if err := uh.uu.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &userv1.UpdateUserResponse{
		Status: userv1.Status_SUCCESS,
	}, nil
}


