package cli

import (
	"context"
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListChallengerDid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-challenger-did",
		Short: "list all challenger did",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllChallengerDidRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ChallengerDidAll(context.Background(), params)
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

func CmdShowChallengerDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-challenger-did [id]",
		Short: "shows a challenger did",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argId := args[0]

			params := &types.QueryGetChallengerDidRequest{
				Id: argId,
			}

			res, err := queryClient.ChallengerDidDocument(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
