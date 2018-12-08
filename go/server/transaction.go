package server

import (
	"encoding/json"
	"net/http"
)

func (a API) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var res APIResponse

	ctx := r.Context()

	txs, err := a.Store.GetAllTransactions(ctx)
	if err != nil {
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.SetData(txs)
	j, _ := json.MarshalIndent(res, "", "  ")

	w.Write(j)

	return
}
