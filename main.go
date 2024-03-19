package main

import (
	"github.com/fukasawa-ryosuke/serve_streaming_grpc_app/service/http"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// HTTPルーティングの初期化
	http.InitializeUsageRoutes(e)

	// HTTPサーバーを起動
	e.Start(":8080")
}
