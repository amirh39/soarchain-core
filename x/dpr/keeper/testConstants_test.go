/** This file is created for tests. Firstly search what you nee if not find then create a new one for you. */
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
)

func SetupNDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(111111)
		items[i].Creator = strconv.Itoa(i)
		items[i].ClientPubkeys = []string{}
		items[i].IsActive = true
		items[i].Vin = []string{strconv.Itoa(0), strconv.Itoa(1)}
		items[i].PidSupported_1To_20 = true
		items[i].PidSupported_21To_40 = false
		items[i].PidSupported_41To_60 = false
		items[i].LengthOfDpr = 5
	}
	return items
}

func SetupNDifDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(22222)
		items[i].Creator = strconv.Itoa(i)
		items[i].ClientPubkeys = []string{}
		items[i].IsActive = true
		items[i].Vin = []string{strconv.Itoa(0), strconv.Itoa(1)}
		items[i].PidSupported_1To_20 = true
		items[i].PidSupported_21To_40 = false
		items[i].PidSupported_41To_60 = false
		items[i].LengthOfDpr = 3
	}
	return items
}

func SetupNDeactiveDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(5677888)
		items[i].Creator = strconv.Itoa(i)
		items[i].ClientPubkeys = []string{}
		items[i].IsActive = false
		items[i].LengthOfDpr = 1
		items[i].PidSupported_1To_20 = false
		items[i].PidSupported_21To_40 = true
		items[i].PidSupported_41To_60 = false
		items[i].Vin = []string{strconv.Itoa(0)}
	}
	return items
}

func SetupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.DprKeeperWithMocks(t, bankMock)

	dpr.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

const (
	CREATOR = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
)
