package cli

import (
	"context"
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListClientDid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-client-did",
		Short: "list all client did",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllClientDidRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ClientDidAll(context.Background(), params)
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

func CmdShowClientDid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-client-did [address]",
		Short: "shows a client did",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argId := args[0]

			params := &types.QueryGetClientDidRequest{
				Id: argId,
			}

			res, err := queryClient.ClientDid(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
