package p2p

import (
	"dprelay/common/rest"
	"net/http"
	"strings"
)

var p2plist = []string{
	"ab9b355be5c6555ecd617db66d195de3c9445a5e@explorer.darkpool.pro:26656",
	"54781620038c717efb3a8f6cf04383747b46d101@explorer.darkpool.pro:36649",
	"ccbe20d929d94e1d5b4f2f9aa30f0b9db588428e@explorer.darkpool.pro:37649",
}













func GetP2Plist(w http.ResponseWriter, r *http.Request) {
	rest.ResponseText(w, strings.Join(p2plist, ","))
}

