package common

import "net/http"

type EndpointHandler func(w http.ResponseWriter, r *http.Request)
