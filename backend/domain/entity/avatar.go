package entity

import "errors"

type Avatar struct {
	UserID int
	// 各部位に装備しているItemIDを保持する（未装備は 0 とする）
	HatID     int
	ShirtID   int
	JacketID  int
	BottomsID int
	ShoesID   int
}

// 装備のロジック
func (a *Avatar) Equip(category ItemCategory, itemID int) error {
	// 渡されたカテゴリ（部位）に応じて、Avatar自身の状態を書き換える
	switch category {
	case CategoryHAT:
		a.HatID = itemID
	case CategorySHIRT:
		a.ShirtID = itemID
	case CategoryJACKET:
		a.JacketID = itemID
	case CategoryBOTTOMS:
		a.BottomsID = itemID
	case CategorySHOES:
		a.ShoesID = itemID
	default:
		// 存在しない部位を指定された場合はドメインエラー
		return errors.New("domain_error: 不正な装備箇所です")
	}

	return nil
}
