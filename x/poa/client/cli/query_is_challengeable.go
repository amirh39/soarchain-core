package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/soar-robotics/soarchain-core/x/poa/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdIsChallengeable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-challengeable [client-addr]",
		Short: "Query isChallengeable",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqClientAddr := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIsChallengeableRequest{

				ClientAddr: reqClientAddr,
			}

			res, err := queryClient.IsChallengeable(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
