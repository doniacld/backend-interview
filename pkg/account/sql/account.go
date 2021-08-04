package sql

import (
	"context"
	"strings"

	"github.com/gustvision/backend-interview/pkg/account"
)

func (s *Store) Fetch(ctx context.Context, f account.Filter) (account.Account, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, user_id, total `)
	b.WriteString(`FROM account `)
	b.WriteString(`WHERE user_id = $1 ;`)

	row := s.QueryRowContext(ctx, b.String(), []interface{}{
		f.UserID,
	})

	var a account.Account

	if err := row.Scan(
		&a.ID,
		&a.UserID,
		&a.Total,
	); err != nil {
		return account.Account{}, err
	}

	//	row.Err undefined (type *sql.Row has no field or method Err)
	return a, nil
}

func (s *Store) FetchMany(ctx context.Context, f account.Filter, callback func(account.Account) error) ([]account.Account, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, user_id, total `)
	b.WriteString(`FROM account `)
	b.WriteString(`WHERE user_id = $1 ;`)

	rows, err := s.QueryContext(ctx, b.String(), []interface{}{
		f.UserID,
	}...)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	accounts := make([]account.Account, 0)
	for rows.Next() {
		var a account.Account

		if err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Total,
		); err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	return accounts, nil
}


// UpdateTotal updates the total for a given account ID
func (s *Store) UpdateTotal(ctx context.Context, f account.Filter) error {
	b := strings.Builder{}

	b.WriteString(`UPDATE account `)
	b.WriteString(`SET total=$1 `)
	b.WriteString(`WHERE id=$2 `)

	if _, err := s.ExecContext(ctx, b.String(), []interface{}{
		f.ID,
		f.Total,
	}...); err != nil {
		return err
	}

	return nil
}