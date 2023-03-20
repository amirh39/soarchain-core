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
		Use:   "gen-guard [guard-pub-key] [v-2-x-pub-key] [v-2-x-addr] [v-2-x-stake] [v-2-x-ip] [v-2-n-pub-key] [v-2-n-addr] [v-2-n-stake] [v-2-n-ip] [runner-pubkey] [runner-addr] [runner-stake] [runner-ip]",
		Short: "Broadcast message gen-guard",
		Args:  cobra.ExactArgs(13),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGuardPubKey := args[0]
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
