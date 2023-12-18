package wasmbinding_test

import (
	"context"
	"testing"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/x/poa"
	keeper "github.com/amirh39/soarchain-core/x/poa/keeper"
	"github.com/amirh39/soarchain-core/x/poa/testutil"
	"github.com/amirh39/soarchain-core/x/poa/types"

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

const (
	Score = "189"
)

const (
	NotValidAddress = "1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	NotValidndex    = "31c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
)

const (
	Challenger_PubKey       = "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251"
	Challenger_Address      = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score        = "189"
	Challenger_StakedAmount = "2000000000utmotus"
	Challenger_NetEarnings  = "0utmotus"
	Challenger_Type         = "v2n"
)

func CreateNReputation(n int) []types.Reputation {
	items := make([]types.Reputation, n)
	for i := range items {
		items[i].PubKey = Challenger_PubKey
		items[i].Address = Challenger_Address
		items[i].Score = Challenger_Score
		items[i].NetEarnings = Challenger_NetEarnings
		items[i].Type = Challenger_Type
	}
	return items
}
