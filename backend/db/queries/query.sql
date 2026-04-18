-- name: GetAvatar :one
-- ユーザーの現在のアバター状態を取得
SELECT * FROM avatars WHERE user_id = $1;

-- name: SaveAvatar :exec
-- アバター状態の保存（着せ替え）
INSERT INTO avatars (
    user_id, hat_id, shirt_id, jacket_id, bottoms_id, shoes_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) ON CONFLICT (user_id) DO UPDATE SET
    hat_id = EXCLUDED.hat_id,
    shirt_id = EXCLUDED.shirt_id,
    jacket_id = EXCLUDED.jacket_id,
    bottoms_id = EXCLUDED.bottoms_id,
    shoes_id = EXCLUDED.shoes_id;

-- name: GetInventory :many
-- インベントリ（所持アイテム一覧）の取得
SELECT item_id FROM user_items WHERE user_id = $1;

-- name: GetItem :one
-- アイテム詳細の取得
SELECT * FROM items WHERE id = $1;