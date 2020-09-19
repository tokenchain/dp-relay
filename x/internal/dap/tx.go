package dap

import (
	"dprelay/common/conf"
	"dprelay/common/rest"
	"dprelay/x/internal/common"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func SendFundExchange(config *conf.Config) common.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var p ReqExchangeSendFund
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		config_chain := fmt.Sprintf("--chain-id=\"%s\"", config.TomlConfig.ChainID)
		memo := fmt.Sprintf("--memo \"%s\"", p.Memo)
		cmd_l1 := exec.Command(config.TomlConfig.CliBin, "tx", "send", config.TomlConfig.ExchangeWalletAddress, p.Address, memo, config_chain, "--generate-only", ">", "tx.json")
		buf2, err := cmd_l1.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			rest.PostErr(w, err)
			return
		}
		signature := "--validate-signatures"
		sign_from := fmt.Sprintf("--from=\"%s\"", config.TomlConfig.ExchangeWalletKeyName)
		cmd_l2 := exec.Command("yes", config.TomlConfig.AccessKey, "|", config.TomlConfig.CliBin, "tx", "sign", signature, sign_from, "tx_signed.json")
		buf2, err = cmd_l2.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			rest.PostErr(w, err)
			return
		}
		sign_fromc := fmt.Sprintf("--node=\"%s\"", config.TomlConfig.NodeLocal)
		cmd_l3 := exec.Command(config.TomlConfig.CliBin, "tx", "broadcast", sign_fromc, "tx_signed.json")
		// $DCLI tx broadcast --node=$NODE tx_signed.json
		buf2, err = cmd_l3.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			rest.PostErr(w, err)
			return
		}
		jsonRes := simplejson.New()
		jsonRes.Set("res", buf2)
		rest.PostResponse(w, jsonRes)
	}
}

func Transfer(config *conf.Config) common.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var q ReqTransferFund
		err := json.NewDecoder(r.Body).Decode(&q)
		if err != nil {
			fmt.Println(err.Error())
			rest.PostErr(w, err)
			return
		}
		document_path := fmt.Sprintf("%skeys_%s", config.TomlConfig.CustomerKeys, q.FromUserKeyName)
		jsonfile, err := os.OpenFile(document_path, os.O_RDONLY, 0755)
		if err != nil {
			fmt.Println(err.Error())
			rest.PostErr(w, err)
			return
		}
		defer jsonfile.Close()
		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonfile)
		cmd_l1 := exec.Command(config.TomlConfig.CliBin, "tx", "treasury", "send", q.ToUserAddress, q.Amount, string(byteValue), "--yes")
		buf2, err := cmd_l1.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			rest.PostErr(w, err)
			return
		}
		jsonRes := simplejson.New()
		jsonRes.Set("res", buf2)
		rest.PostResponse(w, jsonRes)
	}
}
