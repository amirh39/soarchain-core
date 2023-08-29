package wasmbinding_test

import (
	"context"
	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa"
	keeper "soarchain/x/poa/keeper"
	"soarchain/x/poa/testutil"
	"soarchain/x/poa/types"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func SetupMsgServer(t testing.TB) (keeper.Keeper, context.Context) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.PoaKeeperWithMocks(t, bankMock)
	poa.InitGenesis(ctx, *k, *types.DefaultGenesis())
	context := sdk.WrapSDKContext(ctx)
	return *k, context
}

func SetupClientEntity(n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = Index
		items[i].Address = strconv.Itoa(i)
		items[i].Score = "25"
	}
	return items
}

const (
	Address = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	Index   = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	Type    = "mini"
	Score   = "25"
)

const (
	NotValidAddress = "1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	NotValidndex    = "31c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
)
