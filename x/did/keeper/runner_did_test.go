package keeper_test

import (
	"testing"

	keepertest "github.com/soar-robotics/soarchain-core/testutil/keeper"
	"github.com/soar-robotics/soarchain-core/x/did/types"

	"github.com/stretchr/testify/require"
)

func Test_SetGetRunnerDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	newDid := types.RunnerDid{
		Id:      Did,
		PubKey:  PUBKEY,
		Address: ADDRESS,
	}
	keeper.SetRunnerDid(ctx, newDid)

	got1, found1 := keeper.GetRunnerDid(ctx, ADDRESS)
	require.Equal(t, true, found1)
	require.NotNil(t, got1)

	newDid2 := types.RunnerDid{
		Id:      SecondDid,
		PubKey:  PUBKEY,
		Address: ADDRESS2,
	}
	keeper.SetRunnerDid(ctx, newDid2)

	got2, found2 := keeper.GetRunnerDid(ctx, ADDRESS2)
	require.Equal(t, true, found2)
	require.NotNil(t, got2)

	// Test all DIDs
	resDids := keeper.GetAllRunnerDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}

func Test_GetRunnerDidByPubkey(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	newDid := types.RunnerDid{
		Id:      Did,
		PubKey:  PUBKEY,
		Address: ADDRESS,
	}
	keeper.SetRunnerDid(ctx, newDid)

	got, found := keeper.GetRunnerDidUsingPubKey(ctx, newDid.PubKey)
	t.Log("runner did -->", got, found)
	require.Equal(t, true, found)
	require.NotNil(t, got)

}
