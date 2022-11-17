package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"soarchain/x/poa/types"
)

var _ = strconv.Itoa(0)

func CmdGenGuard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-guard [guard-pub-key] [v-2-x-addr] [v-2-x-stake] [v-2-n-addr] [v-2-n-stake] [runner-addr] [runner-stake]",
		Short: "Broadcast message gen-guard",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGuardPubKey := args[0]
			argV2XAddr := args[1]
			argV2XStake := args[2]
			argV2NAddr := args[3]
			argV2NStake := args[4]
			argRunnerAddr := args[5]
			argRunnerStake := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenGuard(
				clientCtx.GetFromAddress().String(),
				argGuardPubKey,
				argV2XAddr,
				argV2XStake,
				argV2NAddr,
				argV2NStake,
				argRunnerAddr,
				argRunnerStake,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
