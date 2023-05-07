package keeper_test

import (
	"testing"

	params "soarchain/app/params"
	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	v2vRClient = "v2v-rx"
	v2vBClient = "v2v-bx"
	v2nBClient = "v2n-bx"
	runner     = "runner"
)

func createEpoachDataForTestRewards(keeper *keeper.Keeper, ctx sdk.Context) types.EpochData {
	item := types.EpochData{
		TotalEpochs: 30,
		EpochV2VRX:  "2",
		EpochV2VBX:  "3",
		EpochV2NBX:  "4",
		EpochRunner: "5",
	}
	keeper.SetEpochData(ctx, item)
	return item
}

func Test_Updatev2vREpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createEpoachDataForTestRewards(keeper, ctx)
	t.Log("created epoach item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, v2vRClient, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found-found-found", found)
	t.Log("rst-rst-rst", rst)
	r.Equal(err, nil)
}

func Test_Updatev2vBEpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createEpoachDataForTestRewards(keeper, ctx)
	t.Log("created epoach item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, v2vBClient, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found-found-found", found)
	t.Log("rst-rst-rst", rst)
	r.Equal(err, nil)
}

func Test_Updatev2nBEpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createEpoachDataForTestRewards(keeper, ctx)
	t.Log("created epoach item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, v2nBClient, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found-found-found", found)
	t.Log("rst-rst-rst", rst)
	r.Equal(err, nil)
}

func Test_UpdateRunnerEpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createEpoachDataForTestRewards(keeper, ctx)
	t.Log("created epoach item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, runner, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found-found-found", found)
	t.Log("rst-rst-rst", rst)
	r.Equal(err, nil)
}

func Test_UpdateNotValidEpochRewards(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createEpoachDataForTestRewards(keeper, ctx)
	t.Log("created epoach item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, "Not_Valid_Client", earnedCoin)
	assert.NotNil(t, err)
}
