package repository

import (
	"context"

	"backend/domain/entity"
	domainRepo "backend/domain/repository"
	"backend/infrastructure/db"
)

var _ domainRepo.ItemRepository = (*itemRepository)(nil)

type itemRepository struct {
	queries *db.Queries
}

func NewItemRepository(queries *db.Queries) *itemRepository {
	return &itemRepository{
		queries: queries,
	}
}

func (r *itemRepository) FindByID(ctx context.Context, itemID int) (*entity.Item, error) {
	row, err := r.queries.GetItem(ctx, int32(itemID))
	if err != nil {
		return nil, err
	}

	return &entity.Item{
		ID:       int(row.ID),
		Category: entity.ItemCategory(row.Category),
	}, nil
}
