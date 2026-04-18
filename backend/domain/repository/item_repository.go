package repository

import (
	"backend/domain/entity"
	"context"
)

type ItemRepository interface {
	FindByID(ctx context.Context, itemID int) (*entity.Item, error)
}
