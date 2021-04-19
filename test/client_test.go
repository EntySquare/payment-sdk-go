package main

import (
	"context"
	"fmt"
	"github.com/entysquare/payment-sdk-go/lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	creds, err := credentials.NewClientTLSFromFile("../testkeys/server.pem", "entysquare.com")
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	t1 := lib.NewPayApiClient(conn)
	// 模拟请求数据
	// 调用gRPC接口
	r, err := t1.Register(context.TODO(), &lib.RegisterReq{
		Num: "15361445990",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)

}
