package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
)

func SimulateMsgChallengeService(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgChallengeService{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ChallengeService simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ChallengeService simulation not implemented"), nil, nil
	}
}
