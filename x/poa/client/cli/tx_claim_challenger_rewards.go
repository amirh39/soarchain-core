package cli

import (
	"fmt"

	"github.com/amirh39/soarchain-core/x/poa/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdClaimChallengerRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-challenger-rewards [amount]",
		Short: "Broadcast a claimChallengerRewards transaction",
		Args:  cobra.ExactArgs(1),
		RunE:  runClaimChallengerRewardsCmd,
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func runClaimChallengerRewardsCmd(cmd *cobra.Command, args []string) error {
	amount := args[0]

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	msg := types.NewMsgClaimChallengerRewards(clientCtx.GetFromAddress().String(), amount)
	if err := msg.ValidateBasic(); err != nil {
		return fmt.Errorf("failed to validate claimChallengerRewards message: %w", err)
	}

	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
}
