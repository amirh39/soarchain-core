package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMintRewardCoins(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	// Execute the function to be tested
	keeper.MintRewardCoins(ctx)

	// Retrieve the updated epoch data after minting rewards
	updatedEpochData, _ := keeper.GetEpochData(ctx)

	// Validate the results
	require.Equal(t, uint64(0), updatedEpochData.V2VRXtotalChallenges, "V2VRXtotalChallenges should be set to 0 after minting rewards")
	require.Equal(t, uint64(0), updatedEpochData.V2VBXtotalChallenges, "V2VBXtotalChallenges should be set to 0 after minting rewards")
	require.Equal(t, uint64(0), updatedEpochData.V2NBXtotalChallenges, "V2NBXtotalChallenges should be set to 0 after minting rewards")
	require.Equal(t, uint64(0), updatedEpochData.RunnerTotalChallenges, "RunnerTotalChallenges should be set to 0 after minting rewards")
	require.Equal(t, uint64(0), updatedEpochData.ChallengerTotalChallenges, "ChallengerTotalChallenges should be set to 0 after minting rewards")

}
