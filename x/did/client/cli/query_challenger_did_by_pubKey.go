package cli

import (
	"strconv"

	"github.com/amirh39/soarchain-core/x/did/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetChallengerDidByPubKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query_challenger_did_by_pubKey [PubKey]",
		Short: "Query getChallengerDidByPubKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPubKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetChallengerDidByPubKeyRequest{
				Pubkey: reqPubKey,
			}
			res, err := queryClient.GetChallengerDidByPubKey(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
