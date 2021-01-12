package account

import "context"

type Transaction struct {
	ID        string
	Amount    float64
	AccountID string
	CreatedAt int64
}

type FilterTransaction struct {
	AccountID string
}

type StoreTransaction interface {
	InsertTransaction(context.Context, Transaction) error
	FetchManyTransaction(context.Context, FilterTransaction, func(Transaction) error) error
}
