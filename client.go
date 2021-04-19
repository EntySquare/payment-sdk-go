package payment

import (
	"context"
	"entysquare/payment-sdk-go/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type PayClient struct {
	conn   *grpc.ClientConn
	client api.PayApiClient
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
	client := api.NewPayApiClient(conn)
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

	return Account{}, err
}

// Get balance for secret keeper
func (p *PayClient) GetBalance(ctx context.Context, secret string, symbol string) (b string, err error) {
	return "", err
}

// transfer from accounts
func (p *PayClient) Transfer(ctx context.Context, secret string, toNum string, symbol string, amount string) (orderNum string, err error) {
	return "", err
}

// withdraw to address
func (p *PayClient) Withdraw(ctx context.Context, secret string, toAddress string, symbol string, amount string) (orderNum string, err error) {
	return "", err
}

// get bind address for account
func (p *PayClient) GetAddr(ctx context.Context, secret string, symbol string) (address string, err error) {
	return "", err
}

// do contract operations that were supported by payment backend
func (p *PayClient) ContractDo(ctx context.Context, secret string, contractAddress string, params []byte, remark string) (orderNum string, err error) {
	return "", err
}

// seek infos of an returned order
func (p *PayClient) GetOrder(ctx context.Context, secret string, orderNum string) (o Order, err error) {
	return Order{}, err
}
