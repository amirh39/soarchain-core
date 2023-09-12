package cli

import (
	"strconv"

	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGenClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-dpr [pidSupported_1_to_20] [pidSupported_1_to_20] [pidSupported_1_to_20] [vin] [lengthOfDpr]",
		Short: "Broadcast message gen-dpr",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argPidSupport1, _ := strconv.ParseBool(args[0])
			argPidSupport2, _ := strconv.ParseBool(args[1])
			argPidSupport3, _ := strconv.ParseBool(args[2])
			argVin := []string{args[3]}
			argLengthOfDpr, _ := strconv.ParseUint(args[4], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenDpr(
				clientCtx.GetFromAddress().String(),
				argPidSupport1,
				argPidSupport2,
				argPidSupport3,
				argVin,
				argLengthOfDpr,
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
