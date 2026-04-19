"use client";

import Link from "next/link";
import styles from "./page.module.css";
import { useAvatar } from "../hooks/useAvatar";
import { AvatarDisplay } from "../components/AvatarDisplay";
import { EquipButton } from "../components/EquipButton";

export default function Home() {
  const { hatId, message, equipItem } = useAvatar(1);

  return (
    <div className={styles.page}>
      <div className={styles.card}>
        <h1 className={styles.title}>👗 Oshalets</h1>
        
        <AvatarDisplay hatId={hatId} />
        
        <EquipButton 
          itemId={101} 
          label="🧢 赤い帽子(ID:101)を装備する" 
          onClick={equipItem} 
        />
        
        <p className={styles.message}>{message}</p>

        {/* 持ち物一覧ページへのリンク */}
        <Link href="/inventory" className={styles.navLink}>
          📦 持ち物一覧を見る →
        </Link>
      </div>
    </div>
  );
}