package ethclient

import (
	"context"
	"net/http"
)

// EthClient interface for eth client
type EthClient interface {
	GetLatestBlockNumber(ctx context.Context) (int64, error)
	GetTransactionsByBlockNumber(ctx context.Context, blockNumber int64) ([]Transaction, error)
}

// ethClient sends rpc requests to the ethereum
type ethClient struct {
	rpcUrl string
	client *http.Client
}

// New creates new EthClient
func New(rpcUrl string) EthClient {
	return &ethClient{rpcUrl: rpcUrl, client: &http.Client{}}
}

// rpcRequest represents structure for rpc request
type rpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

// rpcError represent structure for rpc execution error
type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
