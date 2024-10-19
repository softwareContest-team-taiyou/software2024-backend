package repository

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type BoxEntity struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"VarChar(45)"`
	IsLock bool `gorm:"bool"`
}

type BoxRepository struct {
	dh DatabaseHandler
}

func NewBoxRepository(dh DatabaseHandler) *BoxRepository {
	return &BoxRepository{dh: dh}
}

func (br *BoxRepository) CreateBox(ctx context.Context, box *domain.Box) error {
	newBox := &BoxEntity{
		ID: box.ID,
		Name: box.Name,
		IsLock: box.IsLock,
	}
	if err := br.dh.Conn(ctx).Table("boxes").Create(newBox).Error; err != nil {
		return err
	}
	return nil
}
