package hdwallet

import (
	"net/http"
)

const (
	Name        = "account_name"
	Email       = "account_email"
	Pass        = "account_pass"
	Private_key = "private_key"
	Index       = "at_index"
	FromIndex   = "from_index"
	ToIndex     = "to_index"
	Mnemonic    = "to_index"
)

type (
	EndpointHandler func(w http.ResponseWriter, r *http.Request)

	ReqMnemonic struct {
		Words string   `json:"keywords" yaml:"keywords"`
		Names []string `json:"names" yaml:"names"`
		From  uint32   `json:"from_index" yaml:"from_index"`
	}
	ReqMnemonicOnly struct {
		Words string `json:"keywords" yaml:"keywords"`
	}
)
