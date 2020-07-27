package common

import (
	"os"
	"path/filepath"

)

// If a new config is created, change some of the default tendermint settings
func interceptLoadConfig() (conf *config.Config, err error) {
	tmpConf := config.DefaultConfig()
	err = viper.Unmarshal(tmpConf)
	if err != nil {
		// TODO: Handle with #870
		panic(err)
	}
	rootDir := tmpConf.RootDir

	// Intercept only if the file doesn't already exist



	appConfigFilePath := filepath.Join(rootDir, "config/app.toml")
	if _, err := os.Stat(appConfigFilePath); os.IsNotExist(err) {
		appConf, _ := config.ParseConfig()
		config.WriteConfigFile(appConfigFilePath, appConf)
	}

	viper.SetConfigName("app")
	err = viper.MergeInConfig()

	return conf, err
}

