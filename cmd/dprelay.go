package main

import(
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagInvCheckPeriod = "inv-check-period"
)

var (
	invCheckPeriod uint
)


func main() {
	cobra.EnableCommandSorting = false



	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "dpd",
		Short:             "dpd Blockchain Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(
		genUtilCli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome),
		genUtilCli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, app.DefaultNodeHome),
		genUtilCli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{}, auth.GenesisAccountIterator{}, app.DefaultNodeHome, app.DefaultCLIHome),
		genUtilCli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics),
		AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome),
		oraclesCli.AddGenesisOracleCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome),
		genUtilCli.MigrateGenesisCmd(ctx, cdc),
	)

	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "DXO", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

