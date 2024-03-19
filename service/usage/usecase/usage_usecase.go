package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "github.com/fukasawa-ryosuke/serve_streaming_grpc_app/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUsageUsecase interface {
	GetDessertStream()
}

type usageUsecase struct{}

func NewUsageUsecase() IUsageUsecase {
	return &usageUsecase{}
}

func (uu *usageUsecase) GetDessertStream() {
	startTime := time.Now() // 関数実行開始時刻を記録

	fmt.Printf("\n\n---- Start GetDessertStream ----\n\n")

	// gRPCサーバーとのコネクションを確立
	address := "localhost:10001"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // コネクションでSSL/TLSを使用しない
		grpc.WithBlock(), // コネクションが確立されるまで待機する
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// gRPCクライアントを生成
	client := pb.NewDessertServiceClient(conn)

  // サーバーストリーミングのレスポンス受信処理
	// 1.クライアントが持つGetDessertStreamメソッドを呼んで、サーバーからレスポンスが送られてくるストリームを取得
  // 2.そのストリームのRecvメソッドを呼ぶことでレスポンスを得る

	stream, err := client.GetDessertStream(context.Background(), &pb.DessertRequest{
		Name: "アップルパイ", // このサンプルではリクエストの値は使わない
		Id:   1,
	})
	if err != nil {
		fmt.Printf("could not get dessert stream: %v", err)
	}

	// 非同期処理
	var wg sync.WaitGroup

	// ディナー処理
	wg.Add(1)
	go func() {
		defer wg.Done()
		dinners := []string{"カレー", "スパゲッティ", "寿司", "ラーメン"}

		for _, dinner := range dinners {
			time.Sleep(1 * time.Second) // 時間待ちを1秒に変更
			fmt.Printf("ディナー: %s\n\n", dinner)
		}
	}()

	// デザート受信
	// gRPCの処理側でSend()されたデータをここで受け取る。（デザートのデータを10回受け取る）
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			dessert, err := stream.Recv()
			// ストリーム処理が終了した場合はエラーが「io.EOF」になる
			// 全てのレスポンスを受け取ったことを確認するために、エラーが「io.EOF」かどうかを確認する
			if err == io.EOF {
				fmt.Println("\n\ngRPCストリーム処理完了")
				break
			}
			if err != nil {
				fmt.Printf("デザートの取得でエラー発生！ : %v", err)
				return
			}

			fmt.Printf("デザート: %s, 説明: %s, ランク: %d\n\n", dessert.Name, dessert.Description)
		}
	}()

	wg.Wait()

	defer fmt.Println("\n\n---- End GetDessertStream ----")

	endTime := time.Now()              // 関数実行終了時刻を記録
	duration := endTime.Sub(startTime) // 実行時間を計算
	fmt.Printf("実行時間: %v\n", duration)
}
