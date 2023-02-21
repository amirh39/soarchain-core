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

func CmdRegisterFactoryKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-factory-key [factory-key]",
		Short: "Broadcast message registerFactoryKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFactoryKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterFactoryKey(
				clientCtx.GetFromAddress().String(),
				argFactoryKey,
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
