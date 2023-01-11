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

func CmdUpdateGuard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-guard [fee] [v-2-x-addr] [v-2-x-stake] [v-2-x-ip] [v-2-n-addr] [v-2-n-stake] [v-2-n-ip] [runner-addr] [runner-stake] [runner-ip]",
		Short: "Broadcast message update-guard",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFee := args[0]
			argV2XAddr := args[1]
			argV2XStake := args[2]
			argV2XIp := args[3]
			argV2NAddr := args[4]
			argV2NStake := args[5]
			argV2NIp := args[6]
			argRunnerAddr := args[7]
			argRunnerStake := args[8]
			argRunnerIp := args[9]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateGuard(
				clientCtx.GetFromAddress().String(),
				argFee,
				argV2XAddr,
				argV2XStake,
				argV2XIp,
				argV2NAddr,
				argV2NStake,
				argV2NIp,
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
