package http

import (
	"fmt"
	"github.com/fukasawa-ryosuke/serve_streaming_grpc_app/service/usage/usecase"
	"log"
	"net"

	dessertGrpc "github.com/fukasawa-ryosuke/serve_streaming_grpc_app/grpc"

	"github.com/labstack/echo/v4"
)

func UsageRoutes(g *echo.Group, handler IUsageHandler) {
	g.GET("/sampleGrpc", handler.SampleGrpc)
}

// HTTPルーティングの初期化
func InitializeUsageRoutes(e *echo.Echo) {
	usageUsecase := usecase.NewUsageUsecase() // usecaseを初期化
	usageHandler := NewUsageHandler(usageUsecase) // handlerを初期化

	usageGroup := e.Group("/usage")

	// ルーティングを登録
	UsageRoutes(usageGroup, usageHandler)

	go startGrpcServer()
}

// gRPCサーバー起動
func startGrpcServer() {
	// gRPCサーバをlocalhost:10001で起動
	lis, err := net.Listen("tcp", "localhost:10001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := dessertGrpc.NewServer()

	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
