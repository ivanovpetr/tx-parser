package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ivanovpetr/tx-parser/pkg/ethutil"
)

// SubscribeRequest represents response structure for Subscribe
type SubscribeRequest struct {
	Address string `json:"address"`
}

// Subscribe http handler for /subscribe adds new subscriber to the parser
func (h Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		InternalServerErrorHandler(w, r)
		return
	}
	var req SubscribeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	addr, err := ethutil.ParseAddressFromString(req.Address)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	err = h.storage.AddSubscriber(addr)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}
