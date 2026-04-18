package usecase

// アバター情報の出力用DTO
type AvatarOutput struct {
	UserID    int
	HatID     int
	ShirtID   int
	JacketID  int
	BottomsID int
	ShoesID   int
}

// インベントリ情報の出力用DTO
type InventoryOutput struct {
	UserID       int
	OwnedItemIDs []int
}
