package database

import (
	"context"

	daily "github.com/rate-engineering/dai-ly/go"
)

func (s Store) GetAllTransactions(ctx context.Context) (txs daily.Transactions, err error) {
	sess := s.NewSession(nil)
	txs = make(daily.Transactions, 0)

	_, err = sess.Select("*").From(TransactionsT).Load(&txs)
	return
}
