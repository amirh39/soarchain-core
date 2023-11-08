package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGenDpr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-dpr [supportedPIDs] [duration] [dprBudget] [maxClientCount] [name]",
		Short: "Broadcast message gen-dpr",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			var supportedPIDs types.SupportedPIDs
			err = json.Unmarshal([]byte(args[0]), &supportedPIDs)
			if err != nil {
				return fmt.Errorf("failed to parse supportedPIDs JSON: %w", err)
			}
			duration, _ := strconv.ParseUint(args[1], 10, 64)
			dprBudget := args[2]
			maxClientCount, _ := strconv.ParseUint(args[3], 10, 64)
			name := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenDpr(
				clientCtx.GetFromAddress().String(),
				supportedPIDs,
				duration,
				dprBudget,
				maxClientCount,
				name,
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
