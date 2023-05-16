package cli

import (
	"soarchain/x/poa/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGenRunner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-runner [RunnerPubKey] [RunnerAddr] [RunnerStake] [RunnerIp]",
		Short: "Broadcast message gen-runner",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argRunnerPubKey := args[0]
			argRunnerAddr := args[1]
			argRunnerStake := args[2]
			argRunnerIp := args[3]
			argRunnerCertificate := args[4]
			argRunnerSignature := args[5]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenRunner(
				clientCtx.GetFromAddress().String(),
				argRunnerPubKey,
				argRunnerAddr,
				argRunnerStake,
				argRunnerIp,
				argRunnerCertificate,
				argRunnerSignature,
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
