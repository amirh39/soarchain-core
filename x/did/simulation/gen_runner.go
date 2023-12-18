package simulation

import (
	"math/rand"

	"github.com/amirh39/soarchain-core/x/did/keeper"
	"github.com/amirh39/soarchain-core/x/did/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgGenRunner(
	ak types.AccountKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgGenRunner{
			Signature:   "",
			Certificate: "",
			Creator:     "",
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "GenClient simulation not implemented"), nil, nil
	}
}
