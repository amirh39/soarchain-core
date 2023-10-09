package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SetGetChallengerDid(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	didDocument, privkey := NewChallengerDidDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetChallengerDid(ctx, *didDocument.Document)

	got, found := keeper.GetChallengerDid(ctx, ADDRESS)
	require.Equal(t, true, found)
	require.NotNil(t, got)

	// Test all DIDs
	resDids := keeper.GetAllChallengerDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 1, len(resDids))
}
