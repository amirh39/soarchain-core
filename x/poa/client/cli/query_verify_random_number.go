package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/soar-robotics/soarchain-core/x/poa/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVerifyRandomNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-random-number [pubkey] [message] [vrv] [proof]",
		Short: "Query verifyRandomNumber",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPubkey := args[0]
			reqMessage := args[1]
			reqVrv := args[2]
			reqProof := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryVerifyRandomNumberRequest{

				Pubkey:  reqPubkey,
				Message: reqMessage,
				Vrv:     reqVrv,
				Proof:   reqProof,
			}

			res, err := queryClient.VerifyRandomNumber(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
