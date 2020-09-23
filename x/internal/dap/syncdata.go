package dap

import (
	"dprelay/common/conf"
	"dprelay/common/rest"
	"dprelay/x/CosMos/BlockSync"
	"dprelay/x/internal/common"
	"net/http"
)

func SyncExchangeDat(config *conf.Config, db *BlockSync.DBOperation) common.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if !config.BaseConfig.SyncTransactions {
			config.BaseConfig.SyncTransactions = true
			go BlockSync.SyncLoop(config, db)
			println("sync server is started")
		}
		rest.ResponseOK(w)
	}
}
