package main

import (
	"context"
	"entysquare/payment-sdk-go/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
)

// 业务实现方法的容器
type PayApiClient struct {
}

// 为server定义 DoMD5 方法 内部处理请求并返回结果
// 参数 (context.Context[固定], *test.Req[相应接口定义的请求参数])
// 返回 (*test.Res[相应接口定义的返回参数，必须用指针], error)
func (t *PayApiClient) FindOrder(tx context.Context, req *api.FindOrderReq) (res *api.FindOrderResp, err error) {
	fmt.Println(req)
	data := api.FindOrderResp{Order: &api.Order{
		OrderNum: "22",
		OrderTyp: "333",
	}}
	return &data, nil
}
func (t *PayApiClient) Transfer(context.Context, *api.TransferReq) (*api.TransferResp, error) {
	return nil, nil
}

//func (t *PayApiClient)FindOrder(context.Context, *config.FindOrderReq) (*config.FindOrderResp, error)          {
//	return nil,nil
//}
func (t *PayApiClient) Withdraw(context.Context, *api.WithdrawReq) (*api.WithdrawResp, error) {
	return nil, nil
}
func (t *PayApiClient) Generate(context.Context, *api.GenAddrReq) (*api.GenAddrResp, error) {
	return nil, nil
}
func (t *PayApiClient) CallContract(context.Context, *api.CallContractReq) (*api.CallContractResp, error) {
	return nil, nil
}

func TestServer(t *testing.T) {
	creds, err := credentials.NewServerTLSFromFile("../testkeys/server.pem", "../testkeys/server.key")
	lis, err := net.Listen("tcp", ":8028") //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds)) //创建gRPC服务
	//config.RegisterWaiterServer(s, &payApiClien{cc})
	api.RegisterPayApiServer(s, &PayApiClient{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
