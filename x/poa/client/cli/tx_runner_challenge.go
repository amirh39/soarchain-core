package cli

import (
	"errors"
	"strconv"

	"soarchain/x/poa/types"

	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRunnerChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runner-challenge [runner-pubkey] [client-pubkey-array] [challenge-result]",
		Short: "Broadcast message runner-challenge",
		Args:  cobra.ExactArgs(3), // Expect 3 arguments
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRunnerPubkey := args[0]
			argClientPubkeys := args[1]
			argChallengeResult := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Split argClientPubkeys into individual client data
			clientData := strings.Split(argClientPubkeys, " ")

			// Create a slice to store *types.ClientPublicKey objects
			clientPubkeys := make([]*types.ClientPublicKey, len(clientData))

			for i, data := range clientData {
				// Split each client data into pubkey and N
				parts := strings.Split(data, ",")
				if len(parts) != 2 {
					return errors.New("invalid client data format")
				}

				pubkey := parts[0]
				numberOfMsgsStr := parts[1]

				// Convert numberOfMsgsStr to int32
				numberOfMsgs, err := strconv.ParseInt(numberOfMsgsStr, 10, 32)
				if err != nil {
					return err
				}

				// Create a *types.ClientPublicKey object
				clientPubkey := types.ClientPublicKey{
					P: pubkey,
					N: int32(numberOfMsgs),
				}
				clientPubkeys[i] = &clientPubkey
			}

			msg := types.NewMsgRunnerChallenge(
				clientCtx.GetFromAddress().String(),
				argRunnerPubkey,
				clientPubkeys,
				argChallengeResult,
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
