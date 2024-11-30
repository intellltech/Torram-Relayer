package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Bitcoin           BitcoinConfig `json:"bitcoin"`
	Torram            TorramConfig  `json:"torram"`
	PollIntervalSeconds int         `json:"poll_interval_seconds"`
}

type BitcoinConfig struct {
	RPCURL      string `json:"rpc_url"`
	RPCUser     string `json:"rpc_user"`
	RPCPassword string `json:"rpc_password"`
}

type TorramConfig struct {
	GRPCURL string `json:"grpc_url"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
