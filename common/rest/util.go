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

func PostResponse(w http.ResponseWriter, json *SimpleJson.Json) {
	payload, err := json.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	res(w, payload)
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
	res(w, jsonBytes)
}

func res(w http.ResponseWriter, jsonBytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
