package database

import (
	"context"

	"github.com/gocraft/dbr"
	daily "github.com/rate-engineering/dai-ly/go"
)

func (s Store) CreateTransaction(request daily.Transaction) (tx daily.Transaction, err error) {
	sess := s.NewSession(nil)

	err = sess.InsertInto(TransactionsT).Columns(TXHashC, SenderC, ReceiverC, DelegateC, FeeAmountC, SendAmountC, TokenC, NonceC).Record(request).
		Returning(TXHashC, SenderC, ReceiverC, DelegateC, FeeAmountC, SendAmountC, TokenC, NonceC).Load(&tx)
	_, err = sess.Select("*").From(TransactionsT).Load(&tx)

	return
}

func (s Store) GetTransaction(TXHash string) (tx daily.Transaction, err error) {
	sess := s.NewSession(nil)

	err = sess.Select("*").From(TransactionsT).Where(dbr.Eq(TXHashC, TXHash)).LoadOne(&tx)

	return
}

func (s Store) GetAllTransactions(ctx context.Context) (txs daily.Transactions, err error) {
	sess := s.NewSession(nil)
	txs = make(daily.Transactions, 0)

	_, err = sess.Select("*").From(TransactionsT).Load(&txs)
	return
}

func (s Store) GetQueuedTransaction(TXHash string) (tx daily.Transaction, err error) {
	sess := s.NewSession(nil)

	err = sess.Select("*").From(TransactionsT).Where(dbr.Eq(TXHashC, daily.StatusQueued)).LoadOne(&tx)

	return
}
