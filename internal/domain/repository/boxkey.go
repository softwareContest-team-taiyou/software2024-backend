package repository

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type BoxKey struct {
	BoxID string `varchar(45)`
	KeyID string `varchar(45)`
}

type BoxKeyRepository struct {
	dh DatabaseHandler
}

func NewBoxKeyRepository(dh DatabaseHandler) *BoxKeyRepository {
	return &BoxKeyRepository{dh: dh}
}

func (bkr *BoxKeyRepository) CreateBoxKey(ctx context.Context, box *domain.Box, key *domain.Key) error {
	newBoxKey := &BoxKey{
		BoxID: box.ID,
		KeyID: key.ID,
	}
	if err := bkr.dh.Conn(ctx).Table("box_keys").Create(newBoxKey).Error; err != nil {
		return err
	}
	return nil
}