package cli

import (
	"soarchain/x/poa/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetRunnerByPubKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-runner-by-PubKey [PubKey]",
		Short: "Query getRunnerByPubKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPubKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetRunnerByPubKeyRequest{

				PubKey: reqPubKey,
			}
			res, err := queryClient.GetRunnerByPubKey(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
