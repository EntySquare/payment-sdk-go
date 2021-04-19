package payment

import (
	"context"
	"entysquare/payment-sdk-go/lib"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type PayServer struct {
	conn   *grpc.Server
	server lib.PayApiServer
}

// build new client with ssl encryption
func ServeSSLPayRpc(certPath string, keyPath string, p lib.PayApiServer) (err error) {
	if certPath == "" {
		certPath = "../testkeys/server.pem"
	}
	if keyPath == "" {
		keyPath = "../testkeys/server.key"
	}
	creds, err := credentials.NewServerTLSFromFile(certPath, keyPath)
	lis, err := net.Listen("tcp", ":8028") //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds)) //创建gRPC服务
	server := &PayServer{
		conn:   s,
		server: p,
	}
	lib.RegisterPayApiServer(s, server.server)
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

// shut server with connection closed
func (p *PayServer) Shut(ctx context.Context) {
	p.conn.GracefulStop()
}

type MockPayServer struct {
}

func (p *MockPayServer) Register(ctx context.Context, r *lib.RegisterReq) (*lib.RegisterResp, error) {
	fmt.Println(r.Num)
	return &lib.RegisterResp{
		Secret: r.Num,
	}, nil

}
func (p *MockPayServer) GetBalance(context.Context, *lib.GetBalanceReq) (*lib.GetBalanceResp, error) {
	return nil, nil
}
func (p *MockPayServer) Transfer(context.Context, *lib.TransferReq) (*lib.TransferResp, error) {
	return nil, nil
}
func (p *MockPayServer) Withdraw(context.Context, *lib.WithdrawReq) (*lib.WithdrawResp, error) {
	return nil, nil
}
func (p *MockPayServer) GetAddr(context.Context, *lib.GetAddrReq) (*lib.GetAddrResp, error) {
	return nil, nil
}
func (p *MockPayServer) ContractDo(context.Context, *lib.ContractDoReq) (*lib.ContractDoResp, error) {
	return nil, nil
}
func (p *MockPayServer) GetOrder(context.Context, *lib.GetOrderReq) (*lib.GetOrderResp, error) {
	return nil, nil
}
