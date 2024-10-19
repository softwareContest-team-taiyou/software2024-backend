package repository

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type KeyEntity struct {
	Id string `gorm:"primaryKey"`
	Name string `gorm:"type VARCHAR(45)"`
}

type KeyRepository struct {
	dh DatabaseHandler
}

func NewKeyRepository(dh DatabaseHandler) *KeyRepository {
	return &KeyRepository{dh: dh}
}

func (kr *KeyRepository) CreateKey(ctx context.Context, key *domain.Key) error {
	newKey := &KeyEntity{
		Id: key.ID,
		Name: key.Name,
	}
	if err := kr.dh.Conn(ctx).Table("keys").Create(newKey).Error; err != nil {
		return err
	}
	return nil
}

