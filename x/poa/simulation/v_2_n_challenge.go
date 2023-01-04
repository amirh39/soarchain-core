package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
)

func SimulateMsgV2NChallenge(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgV2NChallenge{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the V2NChallenge simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "V2NChallenge simulation not implemented"), nil, nil
	}
}
