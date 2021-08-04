package dto

type CreateTransactionReq struct {
	UserID string
	AccountID string
	Amount float64
}

type CreateTransactionResp struct {
}

