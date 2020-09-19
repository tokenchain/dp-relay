package rest

import (
	"encoding/json"
	"fmt"
	SimpleJson "github.com/bitly/go-simplejson"
	"net/http"
)

func writeHeadf(w http.ResponseWriter, code int, format string, i ...interface{}) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(fmt.Sprintf(format, i...)))
}
func writeHead(w http.ResponseWriter, code int, txt string) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(txt))
}
func PostErr(w http.ResponseWriter, jsond error) {
	jsonRes := SimpleJson.New()
	jsonRes.Set("message", jsond.Error())
	jsonRes.Set("res", "internal error")
	payload, err := jsonRes.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	res(w, payload, http.StatusBadRequest)
}
func PostResponse(w http.ResponseWriter, json *SimpleJson.Json) {
	payload, err := json.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	res(w, payload, http.StatusOK)
}
func ResponseOK(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
func ResponseText(w http.ResponseWriter, list string) {
	resTxt(w, []byte(list))
}
func Endpoints(w http.ResponseWriter, allEndPoints []string) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: allEndPoints,
	}
	jsonBytes, err := json.MarshalIndent(endpoints, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res(w, jsonBytes, http.StatusOK)
}
func resTxt(w http.ResponseWriter, jsonBytes []byte) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func res(w http.ResponseWriter, jsonBytes []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonBytes)
}
