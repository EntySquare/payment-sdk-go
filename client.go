package payment

import "context"

type PayClient struct {
}

func NewPaymentClient() (p *PayClient, err error) {
	return &PayClient{}, nil
}

func (p *PayClient) GetBalance(ctx context.Context) (b Balance, err error) {
	return Balance{}, err
}

func (p *PayClient) Pay(ctx context.Context) (o Order, err error) {
	return Order{}, err
}

func (p *PayClient) Withdraw(ctx context.Context) (o Order, err error) {
	return Order{}, err
}

//生成钱包地址
func (p *PayClient) GetAddr(ctx context.Context) (err error) {
	return err
}

func (p *PayClient) GetOrder(ctx context.Context) (o Order, err error) {
	return Order{}, err
}

type PayInterface interface {
	GenAddr()
}
