package conf

import "fmt"

// BaseConfig defines the base configuration for a Tendermint node
type BaseConfig struct {
	RootDir          string `mapstructure:"home"`
	ProxyApp         string `mapstructure:"proxy_app"`
	DBPath           string `mapstructure:"db_dir"`
	LogLevel         string `mapstructure:"log_level"`
	ListenAddr       string `mapstructure:"listen_address"`
	SyncTransactions bool   `mapstructure:"on_sync"`
}

type TomlConfig struct {
	ChainID               string
	DemonBin              string
	CliBin                string
	NativeToken           string
	ExchangeWalletAddress string
	ExchangeWalletKeyName string
	AccessKey             string
	NodeLocal             string
	NODERemote            string
	CustomerKeys          string
	SyncRequest           string
	Debug                 bool
	DB                    Database `toml:"database"`
}

type Database struct {
	Server   string   `toml:"servernode"`
	Password string   `toml:"accesskey"`
	DbName   string   `toml:"db_name"`
	DbUser   string   `toml:"db_user"`
	Coins    []string `toml:"available_coins"`
	Ports    []int
	ConnMax  int `toml:"connection_max"`
	Enabled  bool
}

// Config defines the top level configuration for a Tendermint node
type Config struct {
	// Top level options use an anonymous struct
	BaseConfig `mapstructure:",squash"`
	TomlConfig `mapstructure:"external"`
}

// DefaultConfig returns a default configuration for a Tendermint node
func DefaultConfig() *Config {
	return &Config{
		BaseConfig: DefaultBaseConfig(),
	}
}
func (h *Config) DebugLn(a ...interface{}) {
	if h.TomlConfig.Debug {
		fmt.Println(a...)
	}
}

// DefaultBaseConfig returns a default base configuration for a Tendermint node
func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		ProxyApp:         "tcp://127.0.0.1:26658",
		RootDir:          "./config",
		SyncTransactions: false,
	}
}
