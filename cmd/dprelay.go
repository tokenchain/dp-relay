package main

import(
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagInvCheckPeriod = "inv-check-period"
	flagConfigType         = "config-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagConfigPath         = "config-path"
	flagBBCNetwork         = "bbc-network"
)

var (
	invCheckPeriod uint
)
func initFlags() {
	flag.String(flagConfigPath, "", "config path")
	flag.String(flagConfigType, "", "config type, local or aws")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")
	flag.Int(flagBBCNetwork, int(types.TestNetwork), "bbc chain network type")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {

}

