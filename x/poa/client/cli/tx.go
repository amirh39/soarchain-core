package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"soarchain/x/poa/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdChallengeService())
	cmd.AddCommand(CmdRunnerChallenge())
	cmd.AddCommand(CmdClaimMotusRewards())
	cmd.AddCommand(CmdClaimRunnerRewards())
	cmd.AddCommand(CmdRegisterFactoryKey())
	cmd.AddCommand(CmdClaimChallengerRewards())
	// this line is used by starport scaffolding # 1

	return cmd
}
