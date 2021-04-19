package payment

type Account struct {
	Secret string
}

type Order struct {
	OrderNum string
	// a type of TX/CON
	OrderTyp  string
	OrderInfo interface{}
}

// TX type of OrderInfo
type TxOrder struct {
	FromTyp string
	ToTyp   string
	From    string
	To      string
	Amount  int64
}

// CON type of OrderInfo
type ContractOrder struct {
	FromTyp string
	ToTyp   string
	From    string
	To      string
	Payload []byte
}

// general response
type GeneralResp struct {
	Success bool
	Data    interface{}
}

// generate ADDR
type GenAddrReq struct {
	Typ string
}

// withdraw request
type WithdrawReq struct {
	Typ           string
	Amount        int64
	TargetAddress string
	CallbackURL   string
}

// transfer request
type TransferReq struct {
	TargetTyp   string
	TargetData  string
	Amount      int64
	Typ         string
	CallbackURL string
}

// find order request
type FindOrderReq struct {
	OrderNum string
}

// call contract request
type CallContractReq struct {
	Payload []byte
	Typ     string
}

// response for generate address
type GenAddrResp struct {
	Address string
	Typ     string
}

// response for withdraw
type WithdrawResp struct {
	OrderNum string
}

// response for transfer
type TransferResp struct {
	OrderNum string
}

// response for find order
type FindOrderResp struct {
	Order Order
}

// response for call contract
type CallContractResp struct {
	Payload []byte
}

type Api interface {
	Transfer(req *TransferReq) (res *TransferResp)
	FindOrder(req *FindOrderReq) (res *FindOrderResp)
	Withdraw(req *WithdrawReq) (res *WithdrawResp)
	Generate(req *GenAddrReq) (res *GenAddrResp)
	CallContract(req *CallContractReq) (res *CallContractResp)
}
