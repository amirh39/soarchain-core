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
		Use:   "gen-challenger [ChallengerPubKey] [ChallengerAddr] [ChallengerStake] [ChallengerIp] [ChallengerType]",
		Short: "Broadcast message gen-challenger",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argChallengerPubKey := args[0]
			argChallengerAddr := args[1]
			argChallengerStake := args[2]
			argChallengerIp := args[3]
			argChallengerType := args[4]
			argChallengerCertificate := args[5]
			argChallengerSignature := args[6]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenChallenger(
				clientCtx.GetFromAddress().String(),
				argChallengerPubKey,
				argChallengerAddr,
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
