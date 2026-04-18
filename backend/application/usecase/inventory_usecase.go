package usecase

import (
	"backend/domain/repository"
	"context"
)

type InventoryUseCase interface {
	GetInventory(ctx context.Context, userID int) (*InventoryOutput, error)
}

type inventoryUseCase struct {
	inventoryRepo repository.InventoryRepository // 持ち物を知るため
}

func NewInventoryUseCase(inventoryRepo repository.InventoryRepository) *inventoryUseCase {
	return &inventoryUseCase{
		inventoryRepo: inventoryRepo,
	}
}

// 持ち物一覧の取得
func (u *inventoryUseCase) GetInventory(ctx context.Context, userID int) (*InventoryOutput, error) {

	// 指定されたユーザIDでRepo.findインベントリ取得
	// inventoryRepo.Find(userID)で指定されたユーザIDにてInventoryテーブルの特定dbレコードをSELECT
	inventory, err := u.inventoryRepo.Find(ctx, userID)
	if err != nil {
		return nil, err
	}
	// DTOに詰めて返す
	return &InventoryOutput{
		UserID:       inventory.UserID,
		OwnedItemIDs: inventory.OwnedItemIDs,
	}, nil

}
