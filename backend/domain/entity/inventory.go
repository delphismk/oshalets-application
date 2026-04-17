package entity

import "errors"

// 持ち物集約
type Inventory struct {
	UserID       int
	OwnedItemIDs []int
}

// ユーザがそのアイテムを持つかを管理するドメインルール
func (i *Inventory) CheckItem(itemId int) error {
	// Inventory内に指定されたitemIDがあるか
	for _, OwnedItemId := range i.OwnedItemIDs {
		// 一致すればnilを返す
		if OwnedItemId == itemId {
			return nil
		}
	}
	// 一致せずloop抜けてしまったらdomain error
	return errors.New("domain_error: その服は持っていません")
}
