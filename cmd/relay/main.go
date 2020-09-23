package main

import (
	"dprelay/common/conf"
	"dprelay/x"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagInvCheckPeriod     = "inv-check-period"
	flagRootDir            = "root"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagListenerAddress    = "port"
	flagStartServer        = "start-rest-server"
)

func printUsage() {
	fmt.Printf("usage: ./relayer %s 'to/my/path' --port 0.0.0.0:1320\n", flagRootDir)
	fmt.Printf("usage: ./relayer %s 0.0.0.0:1320\n", flagListenerAddress)
}

func initFlags() {
	flag.String(flagListenerAddress, "0.0.0.0:1320", "listening on port")
	flag.String(flagRootDir, "./conf", "the local config dir")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")
	flag.Bool(flagStartServer, false, "starting the rest server")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	er := viper.BindPFlags(pflag.CommandLine)
	if er != nil {
		fmt.Println("there is an error", er.Error())
	}
}

func main() {
	initFlags()
	//printUsage()
	config := conf.DefaultConfig()
	config.RootDir = viper.GetString(flagRootDir)
	config.ListenAddr = viper.GetString(flagListenerAddress)
	if _, err := toml.DecodeFile(fmt.Sprintf("%s/config.toml", config.RootDir), &config.TomlConfig); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The chain ID is now with this %s\n", config.ChainID)
	fmt.Printf("Now this is debug mode: %v\n", config.Debug)
	if viper.GetBool(flagStartServer) {
		fmt.Printf("The DP relay server is now up at %s\n", config.ListenAddr)
		adm := x.NewConf(config)
		adm.Serve()
	}
}
