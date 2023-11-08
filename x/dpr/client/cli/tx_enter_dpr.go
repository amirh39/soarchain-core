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

func CmdEnterDpr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enter-dpr [pubkey] [dprId]",
		Short: "Broadcast message enter-dpr",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argDprId := args[0]
			var argSupportedPIDs types.SupportedPIDs
			err = json.Unmarshal([]byte(args[1]), &argSupportedPIDs)
			if err != nil {
				return fmt.Errorf("failed to parse supportedPIDs JSON: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnterDpr(
				clientCtx.GetFromAddress().String(),
				argDprId,
				argSupportedPIDs,
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
