import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Oshalets App",
  description: "着せ替えアプリ",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}