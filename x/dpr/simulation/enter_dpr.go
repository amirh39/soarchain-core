package simulation

import (
	"math/rand"

	"github.com/amirh39/soarchain-core/x/dpr/keeper"
	"github.com/amirh39/soarchain-core/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgEnterDpr(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEnterDpr{
			Sender: simAccount.Address.String(),
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "EnterDpr simulation not implemented"), nil, nil
	}
}
