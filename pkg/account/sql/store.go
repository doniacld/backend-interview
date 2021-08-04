package sql

import (
	"database/sql"

	"github.com/gustvision/backend-interview/pkg/account"
)

var _ account.Store = (*Store)(nil)

var _ account.StoreTransaction = (*Store)(nil)

type Store struct {
	*sql.DB
}

