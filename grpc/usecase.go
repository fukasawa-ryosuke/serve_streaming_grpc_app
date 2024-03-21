package grpc

import (
	"time"

	pb "github.com/fukasawa-ryosuke/serve_streaming_grpc_app/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DessertStreamServer struct {
	//「サービスの前方互換性を保つために、自作サービス構造体にはこのUnimplementedGreetingServiceServer型を組み込むべき」
	pb.UnimplementedDessertServiceServer
}

// gRPCサーバー作成
func NewServer() *grpc.Server {
	s := grpc.NewServer()

	// gRPCサーバーにDessertStreamServerを登録
	pb.RegisterDessertServiceServer(s, &DessertStreamServer{})

	// gRPCサーバーにリフレクションを登録
  // リフレクションを登録することで、gRPCサーバーに対してgRPCurlなどのツールを使ってリクエストを送信できる
	reflection.Register(s)

	return s
}

// デザート情報をストリームで送信
func (s *DessertStreamServer) GetDessertStream(req *pb.DessertRequest, stream pb.DessertService_GetDessertStreamServer) error {
	desserts := []string{"チーズケーキ", "ティラミス", "マカロン", "エクレア", "カンノーリ", "パンナコッタ", "モンブラン", "クレープ", "シュークリーム", "フルーツタルト"}

	for _, dessertName := range desserts {
		time.Sleep(500 * time.Millisecond)

		// gRPCのストリームを介してクライアントにデータを送信
		// stream.Sendを使うことで呼び出し元の関数にデータを送信することができる
		err := stream.Send(&pb.DessertResponse{
			Description: "美味しい" + dessertName + "です",
			Name:        dessertName,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
