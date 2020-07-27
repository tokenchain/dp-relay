package config

// BaseConfig defines the base configuration for a Tendermint node
type BaseConfig struct {
	ChainID  string
	RootDir  string `mapstructure:"home"`
	ProxyApp string `mapstructure:"proxy_app"`
	DBPath   string `mapstructure:"db_dir"`
	LogLevel string `mapstructure:"log_level"`
}

// Config defines the top level configuration for a Tendermint node
type Config struct {
	// Top level options use an anonymous struct
	BaseConfig `mapstructure:",squash"`
}

// DefaultConfig returns a default configuration for a Tendermint node
func DefaultConfig() *Config {
	return &Config{
		BaseConfig: DefaultBaseConfig(),
	}
}

// DefaultBaseConfig returns a default base configuration for a Tendermint node
func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		ProxyApp: "tcp://127.0.0.1:26658",
		RootDir:  "./config",
	}
}
