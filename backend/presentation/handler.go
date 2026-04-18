package presentation

import (
	"context"

	"connectrpc.com/connect"

	"backend/application/usecase"
	oshaletsv1 "backend/gen/oshalets/v1"
)

// OshaletsHandler は Connect の Handler インターフェースを満たす構造体
type OshaletsHandler struct {
	avatarUsecase    usecase.AvatarUsecase
	inventoryUsecase usecase.InventoryUseCase
}

func NewOshaletsHandler(au usecase.AvatarUsecase, iu usecase.InventoryUseCase) *OshaletsHandler {
	return &OshaletsHandler{
		avatarUsecase:    au,
		inventoryUsecase: iu,
	}
}

// ========================================================
// 1. GetAvatar: アバター状態の取得
// ========================================================
func (h *OshaletsHandler) GetAvatar(
	ctx context.Context,
	req *connect.Request[oshaletsv1.GetAvatarRequest],
) (*connect.Response[oshaletsv1.GetAvatarResponse], error) {

	// ① protoの int32 を Goの int にキャスト
	userID := int(req.Msg.UserId)

	// ② Usecaseの呼び出し
	out, err := h.avatarUsecase.GetAvatar(ctx, userID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// ③ Goの int を protoの int32 にキャストして詰める
	res := connect.NewResponse(&oshaletsv1.GetAvatarResponse{
		UserId:    int32(out.UserID),
		HatId:     int32(out.HatID),
		ShirtId:   int32(out.ShirtID),
		JacketId:  int32(out.JacketID),
		BottomsId: int32(out.BottomsID),
		ShoesId:   int32(out.ShoesID),
	})
	return res, nil
}

// ========================================================
// 2. GetInventory: インベントリ状態の取得
// ========================================================
func (h *OshaletsHandler) GetInventory(
	ctx context.Context,
	req *connect.Request[oshaletsv1.GetInventoryRequest],
) (*connect.Response[oshaletsv1.GetInventoryResponse], error) {

	userID := int(req.Msg.UserId)

	out, err := h.inventoryUsecase.GetInventory(ctx, userID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// ポイント: 配列（スライス）の型変換は for文 を回す必要があります
	var ownedItemIDs []int32
	for _, id := range out.OwnedItemIDs {
		ownedItemIDs = append(ownedItemIDs, int32(id))
	}

	res := connect.NewResponse(&oshaletsv1.GetInventoryResponse{
		UserId:       int32(out.UserID),
		OwnedItemIds: ownedItemIDs,
	})
	return res, nil
}

// ========================================================
// 3. EquipItem: アイテムの着せ替え
// ========================================================
func (h *OshaletsHandler) EquipItem(
	ctx context.Context,
	req *connect.Request[oshaletsv1.EquipItemRequest],
) (*connect.Response[oshaletsv1.EquipItemResponse], error) {

	userID := int(req.Msg.UserId)
	itemID := int(req.Msg.ItemId)

	// EquipItems を呼び出す
	out, err := h.avatarUsecase.EquipItems(ctx, userID, itemID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 着せ替え後の最新状態を返す
	res := connect.NewResponse(&oshaletsv1.EquipItemResponse{
		UserId:    int32(out.UserID),
		HatId:     int32(out.HatID),
		ShirtId:   int32(out.ShirtID),
		JacketId:  int32(out.JacketID),
		BottomsId: int32(out.BottomsID),
		ShoesId:   int32(out.ShoesID),
	})
	return res, nil
}
