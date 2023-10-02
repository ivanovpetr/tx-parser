package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents config for the application
type Config struct {
	Http     HttpConfig     `json:"http"`
	Ethereum EthereumConfig `json:"ethereum"`
	Parser   ParserConfig   `json:"parser"`
}

// HttpConfig represents http config section for the application
type HttpConfig struct {
	Port string `json:"port"`
}

// EthereumConfig represents ethereum config section for the application
type EthereumConfig struct {
	RPCUrl string `json:"rpc_url"`
}

// ParserConfig represents parser config section for the application
type ParserConfig struct {
	LookupInterval int64 `json:"lookup_interval"`
	StartingBlock  int64 `json:"starting_block"`
}

// ReadConfig reads config from path
func ReadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s : %w", path, err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file %s : %w", path, err)
	}

	return &config, nil
}
