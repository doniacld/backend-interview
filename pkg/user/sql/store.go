package sql

import "database/sql"

type Store struct {
	*sql.DB
}
