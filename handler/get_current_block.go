package handler

import (
	"encoding/json"
	"net/http"
)

// GetCurrentBlockResponse represents response structure for GetCurrentBlock
type GetCurrentBlockResponse struct {
	Block int64 `json:"block"`
}

// GetCurrentBlock http handler for /currentBlock endpoint, returns last parsed block
func (h *Handler) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	blockNum, err := h.storage.GetLastParsedBlock()
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
	resp := &GetCurrentBlockResponse{
		Block: blockNum,
	}
	jsonbytes, err := json.Marshal(resp)
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
	w.Write(jsonbytes)
}
