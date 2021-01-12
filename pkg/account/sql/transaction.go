package sql

import (
	"context"
	"strings"

	"github.com/gustvision/backend-interview/pkg/account"
)

func (s *Store) InsertTransaction(ctx context.Context, t account.Transaction) error {
	b := strings.Builder{}

	b.WriteString(`INSERT INTO transaction ( `)
	b.WriteString(`id, amount, account_id, created_at `)
	b.WriteString(`) VALUES ( `)
	b.WriteString(`$1, $2, $3, $4 `)
	b.WriteString(`);`)

	if _, err := s.ExecContext(ctx, b.String(), []interface{}{
		t.ID,
		t.Amount,
		t.AccountID,
		t.CreatedAt,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Store) FetchManyTransaction(
	ctx context.Context,
	f account.FilterTransaction,
	callback func(account.Transaction) error,
) error {
	b := strings.Builder{}
	b.WriteString(`SELECT id, amount, account_id, created_at `)
	b.WriteString(`FROM transaction `)
	b.WriteString(`WHERE account_id = $1 ;`)

	rows, err := s.QueryContext(ctx, b.String(), []interface{}{
		f.AccountID,
	})
	if err != nil {
		return err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var t account.Transaction

		if err := rows.Scan(
			&t.ID,
			&t.Amount,
			&t.AccountID,
			&t.CreatedAt,
		); err != nil {
			return err
		}

		if err := callback(t); err != nil {
			return err
		}
	}

	return rows.Err()
}
