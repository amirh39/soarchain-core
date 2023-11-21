package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/soar-robotics/soarchain-core/x/poa/keeper"
	"github.com/soar-robotics/soarchain-core/x/poa/types"
)

func SimulateMsgSelectRandomRunner(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSelectRandomRunner{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SelectRandomRunner simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SelectRandomRunner simulation not implemented"), nil, nil
	}
}
