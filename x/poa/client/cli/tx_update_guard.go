package cli

import (
	"strconv"

	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateGuard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-guard [v-2-x-pub-key] [v-2-x-addr] [v-2-x-stake] [v-2-x-ip] [v-2-n-pub-key] [v-2-n-addr] [v-2-n-stake] [v-2-n-ip] [runner-pubkey] [runner-addr] [runner-stake] [runner-ip]",
		Short: "Broadcast message update-guard",
		Args:  cobra.ExactArgs(13),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argargV2XPubKey := args[1]
			argV2XAddr := args[2]
			argV2XStake := args[3]
			argV2XIp := args[4]
			argargV2NPubKey := args[5]
			argV2NAddr := args[6]
			argV2NStake := args[7]
			argV2NIp := args[8]
			argRunnerPubKey := args[9]
			argRunnerAddr := args[10]
			argRunnerStake := args[11]
			argRunnerIp := args[12]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateGuard(
				clientCtx.GetFromAddress().String(),
				argargV2XPubKey,
				argV2XAddr,
				argV2XStake,
				argV2XIp,
				argargV2NPubKey,
				argV2NAddr,
				argV2NStake,
				argV2NIp,
				argRunnerPubKey,
				argRunnerAddr,
				argRunnerStake,
				argRunnerIp,
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
