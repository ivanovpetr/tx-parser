package ethclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BlockNumberRPCResponse represents response structure for GetLatestBlockNumber
type BlockNumberRPCResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	Id      int        `json:"id"`
	Result  HexedInt64 `json:"result"`
	Error   *rpcError  `json:"error"`
}

// GetLatestBlockNumber executes rpc eth_blockNumber method
func (c *ethClient) GetLatestBlockNumber(ctx context.Context) (int64, error) {
	payload := rpcRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		Id:      83,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.rpcUrl, bytes.NewBuffer(data))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("rpc request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read rpc response: %w", err)
	}

	var rpcResp BlockNumberRPCResponse
	err = json.Unmarshal(body, &rpcResp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse rpc respone: %w", err)
	}

	if rpcResp.Error != nil {
		return 0, fmt.Errorf("rpc error code: %d message: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	return rpcResp.Result.Value, nil
}
