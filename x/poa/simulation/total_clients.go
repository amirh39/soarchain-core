package simulation

import (
	"math/rand"

	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func SimulateMsgCreateTotalClients(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		msg := &types.MsgCreateTotalClients{
			Creator: simAccount.Address.String(),
		}

		_, found := k.GetTotalClients(ctx)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TotalClients already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateTotalClients(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount          = simtypes.Account{}
			msg                 = &types.MsgUpdateTotalClients{}
			totalClients, found = k.GetTotalClients(ctx)
		)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "totalClients store is empty"), nil, nil
		}
		simAccount, found = FindAccount(accs, totalClients.Creator)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "totalClients creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteTotalClients(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount          = simtypes.Account{}
			msg                 = &types.MsgUpdateTotalClients{}
			totalClients, found = k.GetTotalClients(ctx)
		)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "totalClients store is empty"), nil, nil
		}
		simAccount, found = FindAccount(accs, totalClients.Creator)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "totalClients creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
