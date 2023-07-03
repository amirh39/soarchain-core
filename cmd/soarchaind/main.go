package main

import (
	"os"

	"soarchain/app"
	"soarchain/cmd/soarchaind/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	// optionally print logs, with setting to true logs will print in the console
	os.Setenv("PrintLogs", "true")

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
