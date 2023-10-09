package cli

import (
	"strconv"

	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetRunnerDidByIp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-runner-did-by-ip [ip-address]",
		Short: "Query getRunnerDidByIp",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqIpAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetRunnerDidByIpRequest{
				IpAddress: reqIpAddress,
			}

			res, err := queryClient.GetRunnerDidByIp(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
