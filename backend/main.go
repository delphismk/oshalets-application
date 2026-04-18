package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"backend/application/usecase"
	"backend/gen/oshalets/v1/oshaletsv1connect"
	"backend/infrastructure/db"
	infraRepo "backend/infrastructure/repository"
	"backend/presentation"
)

func main() {
	// =======================================================
	// 1. データベース接続
	// =======================================================
	// Dockerで立てた PostgreSQL への接続情報
	dsn := "postgres://postgres:password@localhost:5432/oshalets?sslmode=disable"
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// DBに繋がっているかPingで確認
	if err := database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// sqlcが生成したクエリクライアントを初期化
	queries := db.New(database)

	// =======================================================
	// 2. DI
	// =======================================================

	// DB層（Infrastructure）はまだ作っていないため、一旦インターフェースにnil
	itemRepo := infraRepo.NewItemRepository(queries)
	avatarRepo := infraRepo.NewAvatarRepository(queries)
	inventoryRepo := infraRepo.NewInventoryRepository(queries)

	// UsecaseにRepoを渡して初期化
	avatarUC := usecase.NewAvatarUseCase(itemRepo, avatarRepo, inventoryRepo)
	inventoryUC := usecase.NewInventoryUseCase(inventoryRepo)

	// HandlerにUsecaseを渡して初期化
	handler := presentation.NewOshaletsHandler(avatarUC, inventoryUC)

	// =======================================================
	// 3. ルーティングの設定 (URLとHandlerを結びつける)
	// =======================================================
	mux := http.NewServeMux()

	// Connectが自動生成した関数を使って、パスとハンドラーを取得
	path, connectHandler := oshaletsv1connect.NewOshaletsServiceHandler(handler)

	// muxに登録（これで "/oshalets.v1.OshaletsService/..." というURLが有効になる）
	mux.Handle(path, connectHandler)

	// (CORS設定)
	// -------------------------------------------------------
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Next.jsからの通信を許可
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"}, // 開発環境なので一旦すべてのヘッダーを許可
	})
	// mux を corsHandler で包む（ラップする）
	corsHandler := c.Handler(mux)

	// =======================================================
	// 4. サーバー起動
	// =======================================================
	address := "localhost:8080"
	fmt.Printf("Server is running on http://%s\n", address)

	// h2c.NewHandler を使うことで、ブラウザ(HTTP/1.1)からの通信も
	// gRPC(HTTP/2)からの通信も、Envoyなしで同時に受けとり
	err = http.ListenAndServe(
		address,
		h2c.NewHandler(corsHandler, &http2.Server{}),
	)

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
