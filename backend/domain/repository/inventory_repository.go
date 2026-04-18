package repository

import (
	"backend/domain/entity"
	"context"
)

type InventoryRepository interface {
	Find(ctx context.Context, userID int) (*entity.Inventory, error)
}
