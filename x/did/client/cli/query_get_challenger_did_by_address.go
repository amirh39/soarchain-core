package cli

import (
	"strconv"

	"github.com/soar-robotics/soarchain-core/x/did/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetChallengerDidByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-challenger-did-by-address [address]",
		Short: "Query getChallengerDidByAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetChallengerDidByAddressRequest{
				Address: reqAddress,
			}

			res, err := queryClient.GetChallengerDidByAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
