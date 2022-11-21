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

func CmdGenGuard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-guard [guard-pub-key] [v-2-x-addr] [v-2-x-stake] [v-2-x-ip] [v-2-n-addr] [v-2-n-stake] [v-2-n-ip] [runner-addr] [runner-stake] [runner-ip]",
		Short: "Broadcast message gen-guard",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGuardPubKey := args[0]
			argV2XAddr := args[1]
			argV2XStake := args[2]
			argV2XIp := args[3]
			argV2NAddr := args[4]
			argV2NStake := args[5]
			argV2NIp := args[6]
			argRunnerAddr := args[7]
			argRunnerStake := args[8]
			argRunnerIp := args[9]

			// string creator = 1;
			// string guardPubKey = 2;
			// string v2XAddr = 3;
			// string v2XStake = 4;
			// string v2XIp = 5;
			// string v2NAddr = 6;
			// string v2NStake = 7;
			// string v2NIp = 8;
			// string runnerAddr = 9;
			// string runnerStake = 10;
			// string runnerIp = 11;

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenGuard(
				clientCtx.GetFromAddress().String(),
				argGuardPubKey,
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
