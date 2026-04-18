package repository

import (
	"backend/domain/entity"
	"context"
)

type AvatarRepository interface {
	Find(ctx context.Context, userID int) (*entity.Avatar, error)
	Save(ctx context.Context, avatar *entity.Avatar) error
}
