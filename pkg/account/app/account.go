package app

import "github.com/gustvision/backend-interview/pkg/account"

type App struct {
	account.Store
	account.StoreTransaction
}
