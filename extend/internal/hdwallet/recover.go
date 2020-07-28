package hdwallet

import (
	"dprelay/common/rest"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"github.com/tokenchain/ixo-blockchain/x/did/exported"
	"net/http"
	"strconv"
)

//(name, mem string, index uint32)
func RecoverySimpleHandler(w http.ResponseWriter, r *http.Request) {
	var p ReqMnemonicOnly
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	index := mux.Vars(r)[Index]
	g, _ := strconv.Atoi(index)
	doc := recover("user_darkpool", p.Words, uint32(g))
	jsonRes := simplejson.New()
	jsonRes.Set("document", doc)
	rest.PostResponse(w, jsonRes)
}
func RecoveryHandler(w http.ResponseWriter, r *http.Request) {
	var p ReqMnemonic
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	list := make([]exported.IxoDid, len(p.Names))
	fromIndex := p.From
	var i uint32
	for i = 1; i < uint32(len(p.Names)); i++ {
		doc := recover(p.Names[i], p.Words, i+fromIndex)
		list = append(list, doc)
	}

	jsonRes := simplejson.New()
	jsonRes.Set("documents", list)
	rest.PostResponse(w, jsonRes)
}
func GenerateMnemonic(w http.ResponseWriter, r *http.Request) {
	generator := exported.NewDidGeneratorBuilder()
	keys := generator.Pre().GetMnemonicString()
	json := simplejson.New()
	json.Set("words", keys)
	rest.PostResponse(w, json)
}
func recover(name, mem string, index uint32) exported.IxoDid {
	account_index := uint32(177)
	generator := exported.NewDidGeneratorBuilder()
	return generator.WithName(name).RecoverBIP44(mem, "", account_index, index)
}
