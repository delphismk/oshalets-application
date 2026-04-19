import { useState, useEffect, useCallback } from "react";
import { client } from "../api/client";

export const useInventory = (userId: number) => {
  const [ownedItemIds, setOwnedItemIds] = useState<number[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  const fetchInventory = useCallback(async () => {
    try {
      setLoading(true);
      // バックエンドの GetInventory を呼び出す
      const res = await client.getInventory({ userId });
      // protoの定義に合わせてレスポンスの配列をセット (例: itemIds)
      setOwnedItemIds(res.ownedItemIds || []); 
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  }, [userId]);

  useEffect(() => {
    fetchInventory();
  }, [fetchInventory]);

  return { ownedItemIds, loading };
};