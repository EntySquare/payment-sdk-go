package main

import (
	"context"
	"entysquare/payment-sdk-go/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	creds, err := credentials.NewClientTLSFromFile("../keys/server.crt", "ygq")
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	t1 := api.NewWaiterClient(conn)
	// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}
	// 调用gRPC接口
	tr, err := t1.DoMD5(context.Background(), &api.Req{JsonStr: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", tr.BackJson)
}
