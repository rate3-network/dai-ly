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

	err = sess.InsertInto(TransactionsT).Columns(TXHashC, SignatureC, SenderC, ReceiverC, DelegateC, FeeAmountC, SendAmountC, TokenC, NonceC, SubmittedHashC).Record(request).
		Returning(TXHashC, SignatureC, SenderC, ReceiverC, DelegateC, FeeAmountC, SendAmountC, TokenC, NonceC, SubmittedHashC, StatusC).Load(&tx)

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

func (s Store) GetAllProcessingTransactions(ctx context.Context) (txs daily.Transactions, err error) {
	sess := s.NewSession(nil)
	txs = make(daily.Transactions, 0)

	_, err = sess.Select("*").From(TransactionsT).Where(dbr.Eq(StatusC, daily.StatusProcessing)).Load(&txs)
	return
}

func (s Store) GetQueuedTransaction(ctx context.Context) (tx daily.Transaction, err error) {
	sess := s.NewSession(nil)

	err = sess.Select("*").From(TransactionsT).Where(dbr.Eq(StatusC, daily.StatusQueued)).
		OrderDir(LastCreatedC, true).LoadOne(&tx)

	return
}

func (s Store) UpdateTransactionStatus(ctx context.Context, TXHash string, submittedHash string, status string) (err error) {
	sess := s.NewSession(nil)

	_, err = sess.Update(TransactionsT).SetMap(map[string]interface{}{
		StatusC:        status,
		SubmittedHashC: submittedHash}).Where(dbr.Eq(TXHashC, TXHash)).Exec()

	return
}

func (s Store) UpdateTransactionStatusBySubmitted(ctx context.Context, submittedHash string, status string) (err error) {
	sess := s.NewSession(nil)

	_, err = sess.Update(TransactionsT).SetMap(map[string]interface{}{
		StatusC:        status,
		SubmittedHashC: submittedHash}).Where(dbr.Eq(SubmittedHashC, submittedHash)).Exec()

	return
}
