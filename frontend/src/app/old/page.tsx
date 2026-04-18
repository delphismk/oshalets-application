"use client"; // 👈 Reactの機能（useState等）を使うための魔法の言葉

import { useEffect, useState } from "react";
import { client } from "../client"; // 先ほど作った通信クライアント

export default function Home() {
  // 画面に表示する状態（State）を定義
  const [hatId, setHatId] = useState<number>(0);
  const [message, setMessage] = useState<string>("アバターを読み込み中...");

  // ① 画面が開いた瞬間に、現在のアバターを取得する
  useEffect(() => {
    const fetchAvatar = async () => {
      try {
        const res = await client.getAvatar({ userId: 1 });
        setHatId(res.hatId);
        setMessage("読み込み完了！");
      } catch (err: any) {
        console.error(err);
        setMessage("エラー: バックエンドが動いていないかも？");
      }
    };
    fetchAvatar();
  }, []);

  // ② ボタンを押した時に、アイテムを装備する関数
  const handleEquipHat = async () => {
      try {
        setMessage("着替え中... 💨");
        // ① バックエンドに「着替える」命令を出す
        await client.equipItem({ userId: 1, itemId: 101 });
        
        // ② 着替え終わったら、「今の姿」をもう一度DBから取得する（確実！）
        const currentAvatar = await client.getAvatar({ userId: 1 });
        setHatId(currentAvatar.hatId);
        
        setMessage("✨ 着替え成功！ ✨");
      } catch (err: any) {
        console.error(err);
        setMessage("エラー: " + err.message);
      }
    };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 text-gray-800">
      <div className="bg-white p-10 rounded-2xl shadow-xl text-center">
        <h1 className="text-3xl font-extrabold mb-8">👗 Oshalets</h1>

        {/* アバター状態の表示エリア */}
        <div className="mb-8 p-6 bg-blue-50 rounded-xl border border-blue-100">
          <p className="text-gray-500 text-sm mb-1">現在の帽子</p>
          <p className="text-5xl font-black text-blue-600">
            {hatId === 0 ? "なし" : `ID: ${hatId}`}
          </p>
        </div>

        {/* 着せ替えボタン */}
        <button
          onClick={handleEquipHat}
          className="px-8 py-4 bg-blue-600 hover:bg-blue-700 text-white font-bold rounded-full transition-all shadow-md active:scale-95"
        >
          🧢 赤い帽子(ID:101)をかぶる！
        </button>

        {/* メッセージ表示エリア */}
        <p className="mt-6 text-sm font-medium text-gray-600 h-4">
          {message}
        </p>
      </div>
    </div>
  );
}