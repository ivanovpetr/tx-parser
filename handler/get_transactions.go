package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/ivanovpetr/tx-parser/pkg/ethclient"
	"github.com/ivanovpetr/tx-parser/pkg/ethutil"
	"github.com/ivanovpetr/tx-parser/pkg/storage"
)

var transactionsForAddress = regexp.MustCompile(`^/transactions/(0x[a-fA-F0-9]+)$`)

// GetTransactionsResponse represents response structure for Transactions
type GetTransactionsResponse struct {
	Transactions []*ethclient.Transaction `json:"transactions"`
}

// Transactions http handler for /transactions/ returns list of transactions for a subscribed address
func (h *Handler) Transactions(w http.ResponseWriter, r *http.Request) {
	matches := transactionsForAddress.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	addr, err := ethutil.ParseAddressFromString(matches[1])
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	txs, err := h.storage.GetAllTransactionsBySubscriber(addr)
	if err != nil {
		if err == storage.NotFoundError {
			NotFoundHandler(w, r)
		} else {
			InternalServerErrorHandler(w, r)
		}
		return
	}

	jsonbytes, err := json.Marshal(GetTransactionsResponse{
		Transactions: txs,
	})
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
	w.Write(jsonbytes)
}
