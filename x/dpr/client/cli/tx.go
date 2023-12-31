package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amirh39/soarchain-core/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/client"
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

	cmd.AddCommand(CmdGenDpr())
	cmd.AddCommand(CmdEnterDpr())
	cmd.AddCommand(CmdLeaveDpr())
	cmd.AddCommand(CmdActivateDpr())
	cmd.AddCommand(CmdUpdateDpr())
	cmd.AddCommand(CmdClaimDprRewards())

	return cmd
}
