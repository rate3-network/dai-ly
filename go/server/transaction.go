package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/gocraft/dbr"
	daily "github.com/rate-engineering/dai-ly/go"
	"github.com/rate-engineering/dai-ly/go/database"
)

const EnvDelegate = "DELEGATE"

func (a API) GetHashHandler(w http.ResponseWriter, r *http.Request) {
	var res APIResponse

	var txReq daily.TransactionRequest
	decodeJSONRequest(&txReq, r)

	hash, err := daily.GenerateHash(txReq)
	if err == database.ErrAlreadyExist {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.SetData(hash)
	j, _ := json.MarshalIndent(res, "", "  ")

	w.Write(j)

	return
}

func (a API) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var res APIResponse

	ctx := r.Context()

	var txReq daily.TransactionRequest
	decodeJSONRequest(&txReq, r)

	hash, err := daily.GenerateHash(txReq)
	if err == database.ErrAlreadyExist {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx := daily.Transaction{
		TXHash:     hash,
		Signature:  txReq.Signature,
		Sender:     txReq.Sender,
		Receiver:   txReq.Receiver,
		Delegate:   os.Getenv(EnvDelegate),
		FeeAmount:  txReq.FeeAmount,
		SendAmount: txReq.SendAmount,
		Token:      txReq.Token,
		Nonce:      txReq.Nonce,
	}

	data, err := a.Store.CreateTransaction(ctx, tx)
	if err != nil && err != dbr.ErrNotFound {
		if err == database.ErrAlreadyExist {
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		a.Logger.With("tx", tx).Error(err)
		res.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.SetData(data)
	j, _ := json.MarshalIndent(res, "", "  ")

	w.Write(j)

	return
}

func (a API) PollTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var res APIResponse
	TXHash := chi.URLParam(r, "hash")

	ctx := r.Context()

	data, err := a.Store.GetTransaction(ctx, TXHash)
	if err != nil && err != dbr.ErrNotFound {
		a.Logger.Error(err)
		res.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err == dbr.ErrNotFound {
		a.Logger.Debug(err)
		res.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	res.SetData(data)
	j, _ := json.MarshalIndent(res, "", "  ")

	w.Write(j)

	return
}

func (a API) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var res APIResponse

	ctx := r.Context()

	data, err := a.Store.GetAllTransactions(ctx)
	if err != nil {
		a.Logger.Error(err)
		res.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.SetData(data)
	j, _ := json.MarshalIndent(res, "", "  ")

	w.Write(j)

	return
}

func (a API) GetOneQueuedTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var res APIResponse

	ctx := r.Context()

	data, err := a.Store.GetQueuedTransaction(ctx)
	if err != nil && err != dbr.ErrNotFound {
		a.Logger.Error(err)
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.SetData(data)
	j, _ := json.MarshalIndent(res, "", "  ")

	w.Write(j)

	return
}
