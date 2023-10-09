package cli

import (
	"soarchain/x/did/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetRunnerDidByPubKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-runner-did-by-PubKey [PubKey]",
		Short: "Query getRunnerDidByPubKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPubKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetRunnerDidByPubKeyRequest{

				Pubkey: reqPubKey,
			}
			res, err := queryClient.GetRunnerDidByPubKey(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
