package repository

import (
	"context"
	"errors"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
	"gorm.io/gorm"
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


func (ur *UserRepository) CheckUserExists(ctx context.Context, id string) (bool, error) {
	// UserEntity のインスタンスを作成します。
	userEntity := &UserEntity{}
	// データベースからユーザーを検索し、結果を userEntity に格納します。
	
	err := ur.dh.Conn(ctx).Table("users").Where("id = ?", id).First(userEntity).Error
	// レコードが見つからない場合は false を返します。
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ユーザーが存在しない場合
			return false, nil
		}
		// その他のエラーが発生した場合
		return false, err
	}
	// ユーザーが存在する場合は true を返します。
	return true, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	// UserEntity のインスタンスを作成します。
	userEntity := &UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
	// データベースのユーザー情報を更新します。
	if err := ur.dh.Conn(ctx).Table("users").Save(userEntity).Error; err != nil {
		return err
	}
	return nil
}