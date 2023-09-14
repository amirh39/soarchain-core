package simulation

import (
	"math/rand"

	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgActivateDpr(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgActivateDpr{
			Sender: simAccount.Address.String(),
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ActivateDpr simulation not implemented"), nil, nil
	}
}
