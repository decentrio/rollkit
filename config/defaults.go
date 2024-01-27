package config

import (
	"time"
)

const (
	// DefaultListenAddress is a default listen address for P2P client.
	DefaultListenAddress = "/ip4/0.0.0.0/tcp/7676"
	// Version is the current rollkit version
	// Please keep updated with each new release
	Version = "0.38.5"
)

// DefaultNodeConfig keeps default values of NodeConfig
var DefaultNodeConfig = NodeConfig{
	P2P: P2PConfig{
		ListenAddress: DefaultListenAddress,
		Seeds:         "",
	},
	Aggregator:     false,
	LazyAggregator: false,
	BlockManagerConfig: BlockManagerConfig{
		BlockTime:   1 * time.Second,
		DABlockTime: 15 * time.Second,
	},
	DAAddress: ":26650",
	Light:     false,
	HeaderConfig: HeaderConfig{
		TrustedHash: "",
	},
}
