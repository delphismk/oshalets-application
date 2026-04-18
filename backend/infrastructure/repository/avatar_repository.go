package repository

import (
	"context"
	"database/sql"

	"backend/domain/entity"
	domainRepo "backend/domain/repository"
	"backend/infrastructure/db"
)

// インターフェースを満たしているかコンパイル時にチェックするGoのおまじない
var _ domainRepo.AvatarRepository = (*avatarRepository)(nil)

type avatarRepository struct {
	queries *db.Queries
}

func NewAvatarRepository(queries *db.Queries) *avatarRepository {
	return &avatarRepository{
		queries: queries,
	}
}

func (r *avatarRepository) Find(ctx context.Context, userID int) (*entity.Avatar, error) {
	// sqlcが生成したメソッドを呼び出す
	row, err := r.queries.GetAvatar(ctx, int64(userID))
	if err != nil {
		if err == sql.ErrNoRows {
			// まだアバターがない場合は、空のアバターを返すかエラーにする
			return &entity.Avatar{UserID: userID}, nil
		}
		return nil, err
	}

	// dbの型からentityの型へマッピング（sqlcのNullInt32をintに変換）
	return &entity.Avatar{
		UserID:    int(row.UserID),
		HatID:     int(row.HatID.Int32),
		ShirtID:   int(row.ShirtID.Int32),
		JacketID:  int(row.JacketID.Int32),
		BottomsID: int(row.BottomsID.Int32),
		ShoesID:   int(row.ShoesID.Int32),
	}, nil
}

func (r *avatarRepository) Save(ctx context.Context, avatar *entity.Avatar) error {
	// entityの型からdbの型へマッピング（0の場合はValid=falseとしてNULL扱いにする）
	params := db.SaveAvatarParams{
		UserID:    int64(avatar.UserID),
		HatID:     sql.NullInt32{Int32: int32(avatar.HatID), Valid: avatar.HatID != 0},
		ShirtID:   sql.NullInt32{Int32: int32(avatar.ShirtID), Valid: avatar.ShirtID != 0},
		JacketID:  sql.NullInt32{Int32: int32(avatar.JacketID), Valid: avatar.JacketID != 0},
		BottomsID: sql.NullInt32{Int32: int32(avatar.BottomsID), Valid: avatar.BottomsID != 0},
		ShoesID:   sql.NullInt32{Int32: int32(avatar.ShoesID), Valid: avatar.ShoesID != 0},
	}

	return r.queries.SaveAvatar(ctx, params)
}
