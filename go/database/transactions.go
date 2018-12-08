package database

import (
	"context"
	"fmt"

	"github.com/gocraft/dbr"
	daily "github.com/rate-engineering/dai-ly/go"
)

var ErrAlreadyExist = fmt.Errorf("Transaction aleady exists")

func (s Store) CreateTransaction(ctx context.Context, request daily.Transaction) (tx daily.Transaction, err error) {
	sess := s.NewSession(nil)

	_, err = s.GetTransaction(ctx, request.TXHash)

	if err != nil && err != dbr.ErrNotFound {
		return
	}

	if err == nil {
		err = ErrAlreadyExist
		return
	}

	err = sess.InsertInto(TransactionsT).Columns(TXHashC, SenderC, ReceiverC, DelegateC, FeeAmountC, SendAmountC, TokenC, NonceC).Record(request).
		Returning(TXHashC, SenderC, ReceiverC, DelegateC, FeeAmountC, SendAmountC, TokenC, NonceC).Load(&tx)

	return
}

func (s Store) GetTransaction(ctx context.Context, TXHash string) (tx daily.Transaction, err error) {
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

func (s Store) GetQueuedTransaction(ctx context.Context) (tx daily.Transaction, err error) {
	sess := s.NewSession(nil)

	err = sess.Select("*").From(TransactionsT).Where(dbr.Eq(TXHashC, daily.StatusQueued)).LoadOne(&tx)

	return
}
