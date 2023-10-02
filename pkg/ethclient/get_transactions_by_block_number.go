package ethclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetBlockByNumberRPCResponse represents reeponse structure for GetTransactionsByBlockNumber
type GetBlockByNumberRPCResponse struct {
	Jsonrpc string    `json:"jsonrpc"`
	Id      int       `json:"id"`
	Result  *Block    `json:"result"`
	Error   *rpcError `json:"error"`
}

// Block represents block with transaction
type Block struct {
	Transactions []Transaction `json:"transactions"`
}

// Transaction represents ethereum transaction
type Transaction struct {
	Hash             string      `json:"hash"`
	From             Address     `json:"from"`
	To               Address     `json:"to"`
	Value            HexedBigInt `json:"value"`
	TransactionIndex HexedInt64  `json:"transactionIndex"`
}

// GetTransactionsByBlockNumber executes rpc eth_getBlockByNumber method
func (c *ethClient) GetTransactionsByBlockNumber(ctx context.Context, blockNumber int64) ([]Transaction, error) {
	payload := rpcRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{fmt.Sprintf("0x%x", blockNumber), true},
		Id:      1,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.rpcUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("rpc request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read rpc response: %w", err)
	}

	var rpcResp GetBlockByNumberRPCResponse
	err = json.Unmarshal(body, &rpcResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rpc respone: %w", err)
	}
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("rpc error code: %d message: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	return rpcResp.Result.Transactions, nil
}
