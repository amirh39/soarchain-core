package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"soarchain/x/poa/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group poa queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListClient())
	cmd.AddCommand(CmdShowClient())
	cmd.AddCommand(CmdListChallenger())
	cmd.AddCommand(CmdShowChallenger())
	cmd.AddCommand(CmdListRunner())
	cmd.AddCommand(CmdShowRunner())
	cmd.AddCommand(CmdListGuard())
	cmd.AddCommand(CmdShowGuard())
	cmd.AddCommand(CmdGetClientByAddress())

	cmd.AddCommand(CmdGetChallengerByAddress())

	cmd.AddCommand(CmdGetRandomChallenger())

	cmd.AddCommand(CmdGetRandomRunner())

	cmd.AddCommand(CmdListVrfData())
	cmd.AddCommand(CmdShowVrfData())
	cmd.AddCommand(CmdListVrfUser())
	cmd.AddCommand(CmdShowVrfUser())
	// this line is used by starport scaffolding # 1

	return cmd
}
