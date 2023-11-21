package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/soar-robotics/soarchain-core/x/soarmint/types"
	"github.com/spf13/cobra"
)

func CmdShowMinter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-minter",
		Short: "shows minter",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetMinterRequest{}

			res, err := queryClient.Minter(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
