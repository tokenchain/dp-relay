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
)

func printUsage() {
	fmt.Print("usage: ./relayer --root-dir 'full path' --port 0.0.0.0:1320\n")
}

func initFlags() {
	flag.String(flagListenerAddress, "0.0.0.0:1320", "listening on port")
	flag.String(flagRootDir, "./conf", "the local config dir")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	er := viper.BindPFlags(pflag.CommandLine)
	if er != nil {
		fmt.Println("there is an error", er.Error())
	}
}

func main() {
	initFlags()
	printUsage()
	config := conf.DefaultConfig()
	config.RootDir = viper.GetString(flagRootDir)
	config.ListenAddr = viper.GetString(flagListenerAddress)
	if _, err := toml.DecodeFile(fmt.Sprintf("%s/config.toml", config.RootDir), &config.TomlConfig); err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.ChainID)
	adm := x.NewConf(config)
	adm.Serve()
}
