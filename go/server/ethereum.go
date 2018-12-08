package server

import (
	"context"
	"fmt"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gocraft/dbr"
	daily "github.com/rate-engineering/dai-ly/go"
)

func (a API) FindAndSubmitQueuedTransaction() {
	ctx := context.Background()

	tx, err := a.Store.GetQueuedTransaction(ctx)
	if err != nil && err != dbr.ErrNotFound {
		a.Logger.Error(err)
		return
	}

	if err == dbr.ErrNotFound {
		a.Logger.Info("No queued transaction found")
		return
	}

	a.Logger.With("queued", tx).Info("Submitting queued transaction")

	submittedHash, err := a.EthereumClient.ExecuteTransaction(tx)
	if err != nil && err != dbr.ErrNotFound {
		a.Logger.Error(err)
		return
	}

	a.Logger.With("submittedHash", submittedHash).Info("Submitted transaction successfully")

	err = a.Store.UpdateTransactionStatus(ctx, tx.TXHash, submittedHash, daily.StatusProcessing)
	if err != nil && err != dbr.ErrNotFound {
		a.Logger.Error(err)
		return
	}

}

func (a API) ListenEvents() {
	ctx := context.Background()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{a.EthereumClient.TokenAddress},
	}

	logs := make(chan types.Log)
	sub, err := a.EthereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		a.Logger.Error(err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			a.Logger.Error(err)
		case vLog := <-logs:
			a.Logger.With("vLog", vLog).Info("Event log found.") // pointer to event log
			// contractAbi, err := abi.JSON(strings.NewReader(string(contract.DelegatableDaiABI)))
			// if err != nil {
			// 	a.Logger.Error(err)
			// 	return
			// }

			err = a.Store.UpdateTransactionStatusBySubmitted(ctx, vLog.TxHash.Hex(), daily.StatusSuccess)
			if err != nil && err != dbr.ErrNotFound {
				a.Logger.Error(err)
				return
			}
		}
	}
}

func (a API) IterateProcessingTransactions() {
	ctx := context.Background()

	txs, err := a.Store.GetAllProcessingTransactions(ctx)
	if err != nil && err != dbr.ErrNotFound {
		a.Logger.Error(err)
		return
	}

	if len(txs) <= 0 {
		a.Logger.Info("No processing transactions found")
	}

	for _, tx := range txs {
		txHash := common.HexToHash(tx.SubmittedHash)
		submittedTx, isPending, err := a.EthereumClient.TransactionByHash(context.Background(), txHash)
		if err != nil {
			a.Logger.With("submittedHash", txHash).Error(err)
			if err.Error() == "not found" {
				now := time.Now()
				now.Add(time.Minute * -2)
				if now.After(tx.LastModified) {
					a.Logger.Info("Re-queuing transaction that was missing for 2 minutes: " + tx.SubmittedHash)
					// retry if processing transaction not found after 2 minutes
					err = a.Store.UpdateTransactionStatusBySubmitted(ctx, tx.SubmittedHash, daily.StatusQueued)
					if err != nil {
						a.Logger.Error(err)
					}

				}
			}
			continue
		}

		if isPending {
			a.Logger.Info("Transaction is still pending: " + submittedTx.Hash().Hex())
		} else {
			receipt, err := a.EthereumClient.TransactionReceipt(context.Background(), txHash)
			if err != nil {
				a.Logger.With("submittedHash", txHash).Error(err)
				continue
			}

			fmt.Println(receipt.Status) // 1
			if receipt.Status == 1 {    // 1 is success
				err = a.Store.UpdateTransactionStatusBySubmitted(ctx, submittedTx.Hash().Hex(), daily.StatusSuccess)
				if err != nil && err != dbr.ErrNotFound {
					a.Logger.Error(err)
					return
				}
				a.Logger.Info("Transaction updated to success: " + submittedTx.Hash().Hex())
			} else {
				err = a.Store.UpdateTransactionStatusBySubmitted(ctx, submittedTx.Hash().Hex(), daily.StatusFailed)
				if err != nil && err != dbr.ErrNotFound {
					a.Logger.Error(err)
					return
				}
				a.Logger.Info("Transaction updated to failed: " + submittedTx.Hash().Hex())
			}

		}

	}
}
