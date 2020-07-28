package hdwallet

import (
	"encoding/json"
	"net/http"
)

func ResponseOK() EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
