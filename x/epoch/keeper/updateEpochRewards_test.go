package keeper_test

import (
	"testing"

	params "github.com/amirh39/soarchain-core/app/params"
	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/x/epoch/constants"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Updatev2vREpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.EpochKeeper(t)
	epoch := CreateEpochData(keeper, ctx)
	t.Log("created epoch item in store", epoch)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, constants.V2VRX, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found---->", found)
	t.Log("rst---->", rst)
	r.Equal(err, nil)
}

func Test_Updatev2vBEpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.EpochKeeper(t)
	item := CreateEpochData(keeper, ctx)
	t.Log("created epoch item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, constants.V2VBX, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found-found-found", found)
	t.Log("rst-rst-rst", rst)
	r.Equal(err, nil)
}

// invalid decimal coin expression

func Test_Updatev2nBEpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.EpochKeeper(t)
	item := CreateEpochData(keeper, ctx)
	t.Log("created epoch item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, constants.V2NBX, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found-found-found", found)
	t.Log("rst-rst-rst", rst)
	r.Equal(err, nil)
}

func Test_UpdateRunnerEpochRewards(t *testing.T) {
	r := require.New(t)
	keeper, ctx := keepertest.EpochKeeper(t)
	item := CreateEpochData(keeper, ctx)
	t.Log("created epoch item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, constants.Runner, earnedCoin)
	rst, found := keeper.GetEpochData(ctx)
	t.Log("found-found-found", found)
	t.Log("rst-rst-rst", rst)
	r.Equal(err, nil)
}

func Test_UpdateNotValidEpochRewards(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	item := CreateEpochData(keeper, ctx)
	t.Log("created epoch item in store", item)
	earnedRewardsInt := sdk.NewIntFromUint64((uint64(23)))
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)
	err := keeper.UpdateEpochRewards(ctx, "Not_Valid_Client", earnedCoin)
	assert.NotNil(t, err)
}
