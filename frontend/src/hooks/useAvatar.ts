import { useState, useEffect, useCallback, use } from "react";
import { client } from "../api/client";

export const useAvatar = (userId: number) => {
    const [hatId, setHatId] = useState<number>(0);
    const [message, setMessage] = useState<string>("読み込み中...");

    // アバター取得のロジック
    const fetchAvatar = useCallback(async () => {
        try {
            const res = await client.getAvatar({ userId });
            setHatId(res.hatId);
            // TODO: 他のitemのセットも必要そう
        } catch (err: any) {
            console.error(err);
            setMessage("エラーが発生しました");
        }
    }, [userId]);

    // 着替えロジック
    const equipItem = async (itemId: number) => {
        try {
            setMessage("着替え中...");
            await client.equipItem({ userId, itemId });
            await fetchAvatar(); // 着替え後に再取得
            setMessage("着替え完了！")
        } catch (err: any) {
            setMessage("エラー: " + err.Message);
        }
    };

    // 初回マウント時に取得
    useEffect(() => {
        fetchAvatar();
    }, [fetchAvatar]);

    // UI componentで使いたいものだけ返す
    return {
        hatId,
        message,
        equipItem,
    };
};