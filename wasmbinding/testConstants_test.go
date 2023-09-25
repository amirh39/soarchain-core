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

func SetupClientEntity(n int) []types.Reputation {
	items := make([]types.Reputation, n)
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
	Score   = "189"
)

const (
	NotValidAddress = "1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	NotValidndex    = "31c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
)

const (
	Challenger_PubKey        = "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251"
	Challenger_Address       = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score         = "189"
	Challenger_StakedAmount  = "2000000000utmotus"
	Challenger_NetEarnings   = "0utmotus"
	Challenger_StakedAmount2 = "2000000000udmotus"
	Challenger_NetEarnings2  = "0udmotus"
	Challenger_IpAddr        = ""
	Challenger_IPAddress     = "104.248.142.45"
	Challenger_Type          = "v2n"
	Challenger_Creator       = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score2        = "82"
)

func CreateNChallenger(n int) []types.Challenger {
	items := make([]types.Challenger, n)
	for i := range items {
		items[i].PubKey = Challenger_PubKey
		items[i].Address = Challenger_Address
		items[i].Score = Challenger_Score
		items[i].StakedAmount = Challenger_StakedAmount
		items[i].NetEarnings = Challenger_NetEarnings
		items[i].IpAddress = ""
		items[i].Type = Challenger_Type

	}
	return items

}
