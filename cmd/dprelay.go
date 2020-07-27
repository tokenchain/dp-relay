package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"main.go/common/c"
	"main.go/extend"
)

const (
	flagInvCheckPeriod     = "inv-check-period"
	flagRootDir            = "root-dir"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagListenerAddress    = "config-path"
)

var (
	invCheckPeriod uint
)

func printUsage() {
	fmt.Print("usage: ./relayer --bbc-network [0 for testnet, 1 for mainnet] --config-path config_file_path\n")
}

func initFlags() {
	flag.String(flagListenerAddress, "0.0.0.0:8080", "listening on port")
	flag.String(flagRootDir, "./conf", "the local config dir")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	initFlags()
	config := c.DefaultConfig()
	config.RootDir = viper.GetString(flagRootDir)
	config.ListenAddr = viper.GetString(flagListenerAddress)
	adm := extend.NewConf(config)
	go adm.Serve()

	select {}
}
