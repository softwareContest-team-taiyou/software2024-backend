package repository

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type UserEntity struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

type UserRepository struct {
	dh DatabaseHandler
}

func NewUserRepository(databaseHandler DatabaseHandler) *UserRepository {
	return &UserRepository{dh: databaseHandler}
}

func (ur *UserRepository) GetUser(ctx context.Context,id string) (*domain.User, error) {
	userEntity := &UserEntity{}
	if err := ur.dh.Conn(ctx).Table("users").Where("id = ?", id).First(userEntity).Error; err != nil {
		return nil, err
	}
	return &domain.User{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	newUser := &UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
	if err := ur.dh.Conn(ctx).Table("users").Create(newUser).Error; err != nil {
		return err
	}
	return nil
}