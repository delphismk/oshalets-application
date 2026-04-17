package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"context"
)

type AvatarUsecase interface {
	GetAvatar(ctx context.Context, input *GetAvatarInput) (*entity.Avatar, nil)
}

type avatarUseCase struct {
	avatarRepo repository.AvatarRepository
}

func NewAvatarUseCase(ar repository.AvatarRepository) *AvatarUseCase {
	return &AvatarUseCase{
		avatarRepo: ar,
	}
}

// アバター状態の取得
func (u *avatarUseCase) GetAvatar(ctx context.Context, userID int) (*entity.Avatar, error) {
	// AvatarのRepo.FindByID(userID)で指定されたidにてAvatarテーブルの特定dbレコードをSELECT

	// entityにマッピングして返す
}

// 指定したユーザに対し指定したアイテムIDを着せる操作
func (u *avatarUseCase) EquipItems(ctx context.Context, userID int, itemID int) (*entity.Avatar, error) {
	// avatarRepo.Find(userID)でavatar持ってくる

	// そもそもそのアイテムを持ってるかのCHK：inventoryのCheckItem(itemId)
	//or
	// inventoryRepo.Find(userID)の返り値で持ってるか持ってないか判断？もうよくわからない

	// ユーザが指定したitemIDのCategoryを取得

	// 装備していい場所かのCHK, avatarのdomainlogicを通すEquip(category, itemID)

	// entityにマッピングして返す

}
