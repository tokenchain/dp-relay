package dap

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
	ReqMnemonic struct {
		Words string   `json:"keywords" yaml:"keywords"`
		Names []string `json:"names" yaml:"names"`
		From  uint32   `json:"from_index" yaml:"from_index"`
	}
	ReqExchangeSendFund struct {
		User    string `json:"username" yaml:"username"`
		Address string `json:"address" yaml:"address"`
		Coin    string `json:"coinname" yaml:"coinname"`
		Amount  string `json:"amount" yaml:"amount"`
		Memo    string `json:"memo" yaml:"memo"`
	}
	ReqTransferFund struct {
		FromUserKeyName string `json:"from_user_name" yaml:"from_user_name"`
		ToUserAddress   string `json:"to_user" yaml:"to_user"`
		Coin            string `json:"coinname" yaml:"coinname"`
		Amount          string `json:"amount" yaml:"amount"`
		Memo            string `json:"memo" yaml:"memo"`
	}
)
