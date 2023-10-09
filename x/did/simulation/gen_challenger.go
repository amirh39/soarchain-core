package simulation

import (
	"math/rand"

	"soarchain/x/did/keeper"
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgGenChallenger(
	ak types.AccountKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgGenChallenger{
			Document:    nil,
			Signature:   "",
			Certificate: "",
			Creator:     "",
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "GenChallenger simulation not implemented"), nil, nil
	}
}
