package x

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tokenchain/ixo-blockchain/app"
	"strings"
)

const (
	certPath             = "/Users/hesk/Documents/ixo/dp-hub/private/did_mainnet"
	appName              = "Darkpool"
	CoinType             = 177
	Bech32MainPrefix     = "dx0"
	Bech32PrefixAccAddr  = Bech32MainPrefix
	Bech32PrefixAccPub   = Bech32MainPrefix + sdk.PrefixPublic
	Bech32PrefixValAddr  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	Bech32PrefixValPub   = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	Bech32PrefixConsAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	Bech32PrefixConsPub  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

func setPrefix() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
	config.Seal()
}

func makeFile(key string, data []byte) {
	filename := fmt.Sprintf("did_%s.json", strings.ToLower(key))
	NewCertWriter(certPath, filename).NewFile(data)
}
func appendImportCert(drive WriterCert, key string) {
	line1 := fmt.Sprintf(
		"ADDRESS_%s=$(jq '.dp.address' $DID_FOLDER/did_%s.json -r)",
		strings.ToUpper(key),
		strings.ToLower(key),
	)

	line2 := fmt.Sprintf(
		"DIDSOVRIN_%s=$(jq -c . $DID_FOLDER/did_%s.json)",
		strings.ToUpper(key),
		strings.ToLower(key),
	)

	line3 := fmt.Sprintf(
		"DID_%s=$(jq '.did' $DID_FOLDER/did_%s.json -r)",
		strings.ToUpper(key),
		strings.ToLower(key),
	)
	drive.AppendLine(line1)
	drive.AppendLine(line2)
	drive.AppendLine(line3)
}
