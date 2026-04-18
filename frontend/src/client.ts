import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

// 自動生成されたサービスの定義をインポート
import { OshaletsService } from "./gen/oshalets/v1/oshalets_connect";

// 通信方法（トランスポート）の設定
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080", // 👈 バックエンドのURL
});

// これでどこからでもバックエンドを呼び出せる「client」が完成！
export const client = createPromiseClient(OshaletsService, transport);