package payment

import (
	"context"
	"github.com/entysquare/payment-sdk-go/lib"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type PayClient struct {
	conn   *grpc.ClientConn
	client lib.PayApiClient
}

// build new client with ssl encryption
func NewSSLPaymentClient(certPath string, url string) (p *PayClient, err error) {
	if certPath == "" {
		// for integrated test compatible
		certPath = "../testkeys/server.pem"
	}
	if url == "" {
		url = "127.0.0.1:8028"
	}
	creds, err := credentials.NewClientTLSFromFile(certPath, "entysquare.com")
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	client := lib.NewPayApiClient(conn)
	return &PayClient{
		conn:   conn,
		client: client,
	}, nil
}

// shut server with connection closed
func (p *PayClient) Shut(ctx context.Context) (err error) {
	return p.conn.Close()
}

// register new customer with secret returned
func (p *PayClient) Register(ctx context.Context, num string) (account Account, err error) {
	resp, err := p.client.Register(ctx, &lib.RegisterReq{
		Num: num,
	}, nil)
	if err != nil {
		return Account{}, err
	}
	return Account{
		Secret: resp.Secret,
	}, nil
}

// Get balance for secret keeper
func (p *PayClient) GetBalance(ctx context.Context, secret string, symbol string) (b string, err error) {
	resp, err := p.client.GetBalance(ctx, &lib.GetBalanceReq{
		Secret: secret,
		Symbol: symbol,
	})
	if err != nil {
		return "", err
	}
	return resp.Balance, err
}

// transfer from accounts
func (p *PayClient) Transfer(ctx context.Context, secret string, toNum string, symbol string, amount string, callbackUrl string) (orderNum string, err error) {
	resp, err := p.client.Transfer(ctx, &lib.TransferReq{
		Secret: secret,
		ToNum:  toNum,
		Symbol: symbol,
		Amount: amount,
		Url:    callbackUrl,
	})
	if err != nil {
		return "", err
	}
	return resp.OrderNum, err
}

// withdraw to address
func (p *PayClient) Withdraw(ctx context.Context, secret string, toAddress string, symbol string, amount string, callbackUrl string) (orderNum string, err error) {
	resp, err := p.client.Withdraw(ctx, &lib.WithdrawReq{
		Secret:    secret,
		ToAddress: toAddress,
		Symbol:    symbol,
		Amount:    amount,
		Url:       callbackUrl,
	})
	if err != nil {
		return "", err
	}
	return resp.OrderNum, err
}

// get bind address for account
func (p *PayClient) GetAddr(ctx context.Context, secret string, symbol string) (address string, err error) {
	resp, err := p.client.GetAddr(ctx, &lib.GetAddrReq{
		Secret:  secret,
		Address: address,
	})
	if err != nil {
		return "", err
	}
	return resp.Address, err
}

// do contract operations that were supported by payment backend
func (p *PayClient) ContractDo(ctx context.Context, secret string, contractAddress string, params []byte, remark string, callbackUrl string) (orderNum string, err error) {
	resp, err := p.client.ContractDo(ctx, &lib.ContractDoReq{
		Secret:          secret,
		ContractAddress: contractAddress,
		Params:          params,
		Remark:          remark,
		Url:             callbackUrl,
	})
	if err != nil {
		return "", err
	}
	return resp.OrderNum, err
}

// seek infos of an returned order
func (p *PayClient) GetOrder(ctx context.Context, secret string, orderNum string) (o Order, err error) {
	resp, err := p.client.GetOrder(ctx, &lib.GetOrderReq{
		Secret:   secret,
		OrderNum: orderNum,
	})
	if err != nil {
		return Order{}, err
	}
	return Order{
		OrderNum:   resp.OrderNum,
		OrderTyp:   resp.OrderTyp,
		State:      resp.State,
		From:       resp.From,
		To:         resp.To,
		CreateTime: resp.CreateTime,
	}, err
}
