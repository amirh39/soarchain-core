package simulation

import (
	"math/rand"

	"soarchain/x/did/keeper"
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgGenDid(
	ak types.AccountKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgGenDid{
			Document:        nil,
			Signature:       nil,
			ClientSignature: "",
			Certificate:     "",
			Creator:         "",
		}

		// TODO: Handling the GenClient simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "GenDid simulation not implemented"), nil, nil
	}
}
