package cli

// import (
// 	"soarchain/x/poa/types"
// 	"strconv"

// 	"github.com/spf13/cobra"
// )

// var _ = strconv.Itoa(0)

// func CmdGenRunner() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "gen-runner [RunnerPubKey] [RunnerAddr] [RunnerStake] [RunnerIp]",
// 		Short: "Broadcast message gen-runner",
// 		Args:  cobra.ExactArgs(4),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			argRunnerPubKey := args[0]
// 			argRunnerAddr := args[1]
// 			argRunnerStake := args[2]
// 			argRunnerIp := args[3]
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgGen
// 		}
// 	}

// }
