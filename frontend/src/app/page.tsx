"use client";

import { useEffect, useState } from "react";
import { client } from "../client";

// IDと画像の対応表（マッピング）
const ITEM_IMAGES: Record<number, string> = {
  0: "https://placehold.jp/24/cccccc/ffffff/200x200.png?text=Empty", // 何もない時
  101: "https://t3.ftcdn.net/jpg/00/61/31/54/360_F_61315414_9O8Z8VvP3HIsL4VjAn2I8K1X8lI6p9W9.jpg", // 帽子の画像例
};

export default function Home() {
  const [hatId, setHatId] = useState<number>(0);
  const [message, setMessage] = useState<string>("読み込み中...");

  // 初期データ取得
  const updateStatus = async () => {
    try {
      const res = await client.getAvatar({ userId: 1 });
      setHatId(res.hatId);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    updateStatus().then(() => setMessage("準備OK！"));
  }, []);

  const handleEquipHat = async () => {
    try {
      setMessage("着替え中...");
      await client.equipItem({ userId: 1, itemId: 101 });
      await updateStatus(); // 着替え終わったら再取得
      setMessage("✨ お似合いですよ！ ✨");
    } catch (err: any) {
      setMessage("エラー: " + err.message);
    }
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-50 p-4">
      <div className="bg-white p-8 rounded-3xl shadow-2xl w-full max-w-md text-center border border-gray-100">
        <h1 className="text-4xl font-black mb-6 text-gray-800 tracking-tight">Oshalets</h1>

        {/* 👗 アバター表示エリア（ここが画像になります！） */}
        <div className="relative w-48 h-48 mx-auto mb-8 bg-gray-100 rounded-2xl overflow-hidden border-4 border-blue-100 flex items-center justify-center">
          <img 
            src={ITEM_IMAGES[hatId] || ITEM_IMAGES[0]} 
            alt="Avatar Hat" 
            className="w-full h-full object-cover transition-all duration-500 transform hover:scale-110"
          />
          {hatId === 0 && <span className="absolute text-gray-400 font-bold">No Item</span>}
        </div>

        <div className="mb-8">
          <p className="text-xs font-bold text-blue-400 uppercase tracking-widest mb-1">Current Hat ID</p>
          <p className="text-3xl font-mono font-black text-gray-700">
            {hatId === 0 ? "NONE" : hatId}
          </p>
        </div>

        <button
          onClick={handleEquipHat}
          className="w-full py-4 bg-blue-600 hover:bg-blue-500 text-white font-bold rounded-2xl transition-all shadow-lg active:transform active:scale-95 mb-4"
        >
          🧢 赤い帽子を装備する
        </button>

        <p className="text-sm font-medium text-gray-400 italic">
          {message}
        </p>
      </div>
    </div>
  );
}