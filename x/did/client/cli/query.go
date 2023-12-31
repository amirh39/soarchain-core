package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/amirh39/soarchain-core/x/did/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group did queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListClientDid())
	cmd.AddCommand(CmdShowClientDid())
	cmd.AddCommand(CmdListRunnerDid())
	cmd.AddCommand(CmdShowRunnerDid())
	cmd.AddCommand(CmdListChallengerDid())
	cmd.AddCommand(CmdShowChallengerDid())
	cmd.AddCommand(CmdGetChallengerDidByAddress())
	cmd.AddCommand(CmdGetChallengerDidByPubKey())
	cmd.AddCommand(CmdGetRunnerDidByPubKey())
	// this line is used by starport scaffolding # 1

	return cmd
}
