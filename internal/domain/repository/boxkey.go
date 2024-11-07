package repository

import (
	"context"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)

type BoxKey struct {
	BoxsID string `varchar(45)`
	KeysID string `varchar(45)`
}

type BoxKeyRepository struct {
	dh DatabaseHandler
}

func NewBoxKeyRepository(dh DatabaseHandler) *BoxKeyRepository {
	return &BoxKeyRepository{dh: dh}
}

func (bkr *BoxKeyRepository) CreateBoxKey(ctx context.Context, box *domain.Box, key *domain.Key) error {
	newBoxKey := &BoxKey{
		BoxsID: box.ID,
		KeysID: key.ID,
	}
	if err := bkr.dh.Conn(ctx).Table("boxs_on_keys").Create(newBoxKey).Error; err != nil {
		return err
	}
	return nil
}
// boxId から keyId を取得する
func (bkr *BoxKeyRepository) GetKeyIdByBox(ctx context.Context, boxId string) (string, error) {
	boxKeyEntity := &BoxKey{}
	if err := bkr.dh.Conn(ctx).Table("boxs_on_keys").Where("boxs_id = ?", boxId).Find(boxKeyEntity).Error; err != nil {
		return "", err
	}
	return boxKeyEntity.KeysID, nil
}