/** This file is created for tests. Firstly search what you need if not find then create a new one for you. */
package keeper_test

import (
	"context"
	"soarchain/x/dpr"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/testutil"
	"soarchain/x/dpr/types"
	"strconv"
	"testing"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	epochKeeper "soarchain/x/epoch/keeper"
	epochTypes "soarchain/x/epoch/types"
)

func CreateDpr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = DprId
		items[i].ClientPubkeys = []string{PUBKEY}
		items[i].Creator = CREATOR
		items[i].Duration = 12
		items[i].IsActive = false
		items[i].Vin = []string{VIN}

		keeper.SetDpr(ctx, items[i])
	}
	return items
}

func SetupDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = DprId
		items[i].Creator = CREATOR
		items[i].Duration = 12
		items[i].IsActive = false

	}
	return items
}

func SetupSecondDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = DprId
		items[i].ClientPubkeys = []string{PUBKEY}
		items[i].Creator = CREATOR
		items[i].Duration = 12
		items[i].IsActive = false
		items[i].Vin = []string{VIN}

	}
	return items
}

func CreateDeactiveDpr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)
		items[i].Creator = strconv.Itoa(i)
		items[i].Duration = 12
		items[i].IsActive = false

		keeper.SetDpr(ctx, items[i])
	}
	return items
}

func CreateAeactiveDpr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)
		items[i].Creator = strconv.Itoa(i)
		items[i].Duration = 12
		items[i].IsActive = true

		keeper.SetDpr(ctx, items[i])
	}
	return items
}

func SetupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.DprKeeperWithMocks(t, bankMock)

	dpr.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

const (
	CREATOR = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	ADDRESS = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	PUBKEY  = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	VIN     = "1HGCM82636c678d14c93ad5bf14448da57f4f241b77e30a013d54f5d76c8126a7029aeb86"
)

var PIDS = []bool{true, false, false}

const (
	Did                  = "did:soar:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm"
	SecondDid            = "did1:soar:1Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgap"
	VerificationMethodId = Did + "#key1"
	DprId                = "uuid-Id"
)

func CreateEpochData(keeper *epochKeeper.Keeper, ctx sdk.Context) epochTypes.EpochData {
	item := epochTypes.EpochData{
		TotalEpochs:                   30,
		EpochV2VRX:                    "2udmotus",
		EpochV2VBX:                    "3udmotus",
		EpochV2NBX:                    "4udmotus",
		EpochRunner:                   "5udmotus",
		EpochChallenger:               "6",
		V2VRXTotalChallenges:          7,
		V2VBXTotalChallenges:          8,
		V2NBXTotalChallenges:          9,
		RunnerTotalChallenges:         10,
		ChallengerTotalChallenges:     11,
		V2VRXLastBlockChallenges:      1,
		V2VBXLastBlockChallenges:      1,
		V2NBXLastBlockChallenges:      1,
		RunnerLastBlockChallenges:     1,
		ChallengerLastBlockChallenges: 1,
		ChallengerPerChallengeValue:   1000000,
		V2NBXPerChallengeValue:        3000000,
		RunnerPerChallengeValue:       1000000,
		InitialPerChallengeValue:      9000000.0,
		TotalChallengesPrevDay:        99,
		V2VBXPerChallengeValue:        3000000,
		V2VRXPerChallengeValue:        3000000,
	}
	keeper.SetEpochData(ctx, item)
	return item
}
