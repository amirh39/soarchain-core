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
		Use:   "update-guard [v-2-x-addr] [v-2-x-stake] [v-2-x-ip] [v-2-n-addr] [v-2-n-stake] [v-2-n-ip] [runner-addr] [runner-stake] [runner-ip]",
		Short: "Broadcast message update-guard",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argV2XAddr := args[0]
			argV2XStake := args[1]
			argV2XIp := args[2]
			argV2NAddr := args[3]
			argV2NStake := args[4]
			argV2NIp := args[5]
			argRunnerAddr := args[6]
			argRunnerStake := args[7]
			argRunnerIp := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateGuard(
				clientCtx.GetFromAddress().String(),
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
