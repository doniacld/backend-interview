package dto

// CreateTransactionReq the request endpoint
type CreateTransactionReq struct {
	AccountID string
	Amount    float64
}

// CreateTransactionResp the response endpoint
type CreateTransactionResp struct {
}
