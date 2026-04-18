package repository

import (
	"context"

	"backend/domain/entity"
	domainRepo "backend/domain/repository"
	"backend/infrastructure/db"
)

var _ domainRepo.InventoryRepository = (*inventoryRepository)(nil)

type inventoryRepository struct {
	queries *db.Queries
}

func NewInventoryRepository(queries *db.Queries) *inventoryRepository {
	return &inventoryRepository{
		queries: queries,
	}
}

func (r *inventoryRepository) Find(ctx context.Context, userID int) (*entity.Inventory, error) {
	// sqlcが生成したメソッドを呼び出す（戻り値は []int32）
	itemIDs, err := r.queries.GetInventory(ctx, int64(userID))
	if err != nil {
		return nil, err
	}

	// []int32 を []int に詰め替える
	var ownedItems []int
	for _, id := range itemIDs {
		ownedItems = append(ownedItems, int(id))
	}

	return &entity.Inventory{
		UserID:       userID,
		OwnedItemIDs: ownedItems,
	}, nil
}
