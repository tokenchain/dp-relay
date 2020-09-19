package x

import (
	"dprelay/common/conf"
	"dprelay/common/rest"
	"dprelay/x/internal/common"
	"dprelay/x/internal/dap"
	"dprelay/x/internal/hdwallet"
	"dprelay/x/internal/p2p"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const (
	DefaultListenAddr = "0.0.0.0:8080"
	Version           = "0.2"
	Commit            = "0.2"
	Name              = "Darkpool Relay"
)


type Centere struct {
	Config     *conf.Config
	routesGet  map[string]common.EndpointHandler
	routesPost map[string]common.EndpointHandler
}

func (center *Centere) Endpoints(w http.ResponseWriter, r *http.Request) {
	var list []string
	for k, _ := range center.routesGet {
		list = append(list, k)
	}
	for k, _ := range center.routesPost {
		list = append(list, k)
	}
	rest.Endpoints(w, list)
}
func NewConf(config *conf.Config) *Centere {
	setPrefix()
	center := &Centere{
		Config: config,
		routesGet: map[string]common.EndpointHandler{
			"/hdwallet/create/mnemonic": hdwallet.GenerateMnemonic,
			"/p2p":                      p2p.GetP2Plist,
			//fmt.Sprintf("/hdwallet/create/{%s}/{%s}/", hdwallet.Name, hdwallet.Index):            hdwallet.RecoveryHandler,
			//fmt.Sprintf("/hdwallet/createbatch/{%s}/{%s}", hdwallet.FromIndex, hdwallet.ToIndex): hdwallet.GenerateMnemonic,
			//"/hdwallet/recover/": hdwallet.GenerateMnemonic,
		},
		routesPost: map[string]common.EndpointHandler{
			fmt.Sprintf("/hdwallet/create/{%s}/", hdwallet.Index): hdwallet.RecoverySimpleHandler,
			"/hdwallet/recovery":   hdwallet.RecoveryHandler,
			"/tx/exchangesend":     dap.SendFundExchange(config),
			"/tx/darkpooltransfer": dap.Transfer(config),
		},
	}
	center.routesGet["/"] = center.Endpoints

	return center
}

func (center *Centere) registerRoutes(r *mux.Router) {
	for k, v := range center.routesGet {
		r.HandleFunc(k, v).Methods("GET")
	}
	for k, v := range center.routesPost {
		r.HandleFunc(k, v).Methods("POST")
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
