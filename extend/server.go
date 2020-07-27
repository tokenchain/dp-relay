package extend

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"main.go/common/c"
	"net/http"
	"time"
)

type (
	Centere struct {
		Config *c.Config
	}
	ERout func(w http.ResponseWriter, r *http.Request)
)

const (
	DefaultListenAddr = "0.0.0.0:8080"
)

var (
	routes = map[string]ERout{
		"/": Endpoints,
		"/hdwallet/create/{at_index}/": SkipSequence,
		"/hdwallet/createbatch/{from_index}/{to_index}": SkipSequence,
		"/hdwallet/recover/": SkipSequence,
	}
)

func getFunctionList() []string {
	var list []string
	for k, _ := range routes {
		list = append(list, k)
	}
	return list
}

func NewConf(config *c.Config) *Centere {
	return &Centere{
		Config: config,
	}
}

func Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: getFunctionList(),
	}
	jsonBytes, err := json.MarshalIndent(endpoints, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (center *Centere) registerRoutes(r *mux.Router) {
	for k, v := range routes {
		r.HandleFunc(k, v)
	}
}
func (center *Centere) Serve() {
	router := mux.NewRouter()
	center.registerRoutes(router)
	listenAddr := DefaultListenAddr
	if center.Config.BaseConfig.ListenAddr != "" {
		listenAddr = center.Config.BaseConfig.ListenAddr
	}
	srv := &http.Server{
		Handler:      router,
		Addr:         listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//util.Logger.Infof("start center server at %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("start center server error, err=%s", err.Error()))
	}
}
