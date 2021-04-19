package main

import (
	"context"
	"entysquare/payment-sdk-go/api"
	"fmt"
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
	t1 := api.NewPayApiClient(conn)
	// 模拟请求数据
	// 调用gRPC接口
	r, err := t1.FindOrder(context.TODO(), &api.FindOrderReq{OrderNum: "213123"})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	//tr, err := t1.DoMD5(context.Background(), &api.Req{JsonStr: res})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("服务端响应: %s", tr.BackJson)
}
