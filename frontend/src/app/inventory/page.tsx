"use client";

import Link from "next/link";
import homeStyles from "../page.module.css";
import styles from "./page.module.css";
import { useInventory } from "../../hooks/useInventory";
import { ITEM_MASTER } from "../../constants/items";

export default function InventoryPage() {
  const { ownedItemIds, loading } = useInventory(1);

  return (
    <div className={homeStyles.page}>
      <div className={homeStyles.card}>
        <h1 className={homeStyles.title}>📦 持ち物一覧</h1>

        {loading ? (
          <p className={homeStyles.message}>読み込み中...</p>
        ) : (
          <ul className={styles.list}>
            {ownedItemIds.length === 0 && <li>何も持っていません</li>}
            {ownedItemIds.map((id) => (
              <li key={id} className={styles.listItem}>
                <img src={ITEM_MASTER[id]?.img} alt="" className={styles.itemImage} />
                <span>{ITEM_MASTER[id]?.name || `不明なアイテム(ID:${id})`}</span>
              </li>
            ))}
          </ul>
        )}

        <Link href="/" className={homeStyles.navLink}>
          ← 着せ替え画面に戻る
        </Link>
      </div>
    </div>
  );
}