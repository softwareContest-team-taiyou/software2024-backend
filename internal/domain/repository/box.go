package repository

import (
	"context"
	"fmt"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type BoxEntity struct {
	ID   string `gorm:"primaryKey"`
	UserId string `gorm:"VsaChar(45)"`
	Name string `gorm:"VarChar(45)"`
	IsLock bool `gorm:"bool"`
}

type BoxRepository struct {
	dh DatabaseHandler
}

func NewBoxRepository(dh DatabaseHandler) *BoxRepository {
	return &BoxRepository{dh: dh}
}

func (br *BoxRepository) CreateBox(ctx context.Context, box *domain.Box,userId  string) error {
	newBox := &BoxEntity{
		ID: box.ID,
		UserId: userId,
		Name: box.Name,
		IsLock: box.IsLock,
	}
	fmt.Printf(newBox.UserId)
	if err := br.dh.Conn(ctx).Table("boxs").Create(newBox).Error; err != nil {
		return err
	}
	return nil
}

func (br *BoxRepository) IsLock(ctx context.Context, UserId string) (*domain.Box, error) {
	boxEntity := &BoxEntity{}
	if err := br.dh.Conn(ctx).Table("boxs").Where("user_id = ?", UserId).First(boxEntity).Error; err != nil {
		return nil, err
	}
	return &domain.Box{
		ID:   boxEntity.ID,
		Name : boxEntity.Name,
		IsLock: boxEntity.IsLock,
	}, nil
}

func (br *BoxRepository) Lock(ctx context.Context, UserId string) error {
	boxEntity := &BoxEntity{}
	if err := br.dh.Conn(ctx).Table("boxs").Where("user_id = ?", UserId).First(boxEntity).Error; err != nil {
		return err
	}
	boxEntity.IsLock = true
	if err := br.dh.Conn(ctx).Table("boxs").Save(boxEntity).Error; err != nil {
		return err
	}
	return nil
}

func (br *BoxRepository) Unlock(ctx context.Context, UserId string) error {
	boxEntity := &BoxEntity{}
	if err := br.dh.Conn(ctx).Table("boxs").Where("user_id = ?", UserId).First(boxEntity).Error; err != nil {
		return err
	}
	boxEntity.IsLock = false
	if err := br.dh.Conn(ctx).Table("boxs").Save(boxEntity).Error; err != nil {
		return err
	}
	return nil
}

