package usecase

import (
	"backend/domain/repository"
	"context"
)

type AvatarUsecase interface {
	GetAvatar(ctx context.Context, userID int) (*AvatarOutput, error)
	EquipItems(ctx context.Context, userID int, itemID int) (*AvatarOutput, error)
}

type avatarUseCase struct {
	itemRepo      repository.ItemRepository      //なんの部位(カテゴリー)かを知るため
	avatarRepo    repository.AvatarRepository    // 着替えさせるため
	inventoryRepo repository.InventoryRepository // 持ち物を知るため
}

func NewAvatarUseCase(itr repository.ItemRepository, ar repository.AvatarRepository, ivr repository.InventoryRepository) AvatarUsecase {
	return &avatarUseCase{
		itemRepo:      itr,
		avatarRepo:    ar,
		inventoryRepo: ivr,
	}
}

// アバター状態の取得
func (u *avatarUseCase) GetAvatar(ctx context.Context, userID int) (*AvatarOutput, error) {
	// AvatarのRepo.FindByID(userID)で指定されたidにてAvatarテーブルの特定dbレコードをSELECT
	avatar, err := u.avatarRepo.Find(ctx, userID)
	if err != nil {
		return nil, err
	}
	// AvatarOutputにマッピングして返す
	return &AvatarOutput{
		UserID:    avatar.UserID,
		HatID:     avatar.HatID,
		ShirtID:   avatar.ShirtID,
		JacketID:  avatar.JacketID,
		BottomsID: avatar.BottomsID,
		ShoesID:   avatar.ShoesID,
	}, nil
}

// 指定したユーザに対し指定したアイテムIDを着せる操作
func (u *avatarUseCase) EquipItems(ctx context.Context, userID int, itemID int) (*AvatarOutput, error) {
	// 指定されたアイテムをユーザが持ってるかCHK
	// inventoryRepo.Find(userID)の返り値でユーザが持つitem一覧(OwnedItemIDs)の取得
	invetory, err := u.inventoryRepo.Find(ctx, userID)
	if err != nil {
		return nil, err
	}
	// inventoryのCheckItem(itemId)で指定されたitemIDがOwnedItemIDの中にあるかCHK
	if err := invetory.CheckItem(itemID); err != nil {
		return nil, err
	}

	// 指定されたItemをアバターに装備する
	// 現状のavatar状態の取得
	// avatarRepo.Find(userID)でavatar持ってくる
	avatar, err := u.avatarRepo.Find(ctx, userID)
	if err != nil {
		return nil, err
	}

	// ユーザが指定したアイテムのCategoryを取得
	// item.Find(itemID)
	item, err := u.itemRepo.FindByID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	// 装備していい場所かのCHK, avatarのdomainlogicを通すEquip(category, itemID)
	if err := avatar.Equip(item.Category, itemID); err != nil {
		return nil, err
	}

	// アバター状態の永続化
	// avatarRepo.Save(avatar)で保存
	if err := u.avatarRepo.Save(ctx, avatar); err != nil {
		return nil, err
	}

	// AvatarOutputにマッピングして返す
	return &AvatarOutput{
		UserID:    avatar.UserID,
		HatID:     avatar.HatID,
		ShirtID:   avatar.ShirtID,
		JacketID:  avatar.JacketID,
		BottomsID: avatar.BottomsID,
		ShoesID:   avatar.ShoesID,
	}, nil

}
