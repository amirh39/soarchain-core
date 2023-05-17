package cli

import (
	"soarchain/x/poa/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGenChallenger() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-challenger [ChallengerStake] [ChallengerIp] [ChallengerType] [hallengerCertificate] [ChallengerSignature]",
		Short: "Broadcast message gen-challenger",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argChallengerStake := args[0]
			argChallengerIp := args[1]
			argChallengerType := args[2]
			argChallengerCertificate := "-----BEGIN CERTIFICATE-----\n" + args[3] + "\n-----END CERTIFICATE-----"
			argChallengerSignature := args[4]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenChallenger(
				clientCtx.GetFromAddress().String(),
				argChallengerStake,
				argChallengerIp,
				argChallengerType,
				argChallengerCertificate,
				argChallengerSignature,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
