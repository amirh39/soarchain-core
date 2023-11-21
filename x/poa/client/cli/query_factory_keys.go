package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/soar-robotics/soarchain-core/x/poa/types"
	"github.com/spf13/cobra"
)

func CmdListFactoryKeys() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-factory-keys",
		Short: "list all factoryKeys",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFactoryKeysRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.FactoryKeysAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowFactoryKeys() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-factory-keys [id]",
		Short: "shows a factoryKeys",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetFactoryKeysRequest{
				Id: id,
			}

			res, err := queryClient.FactoryKeys(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
