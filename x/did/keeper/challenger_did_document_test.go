package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SetGetChallengerDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	didDocument1, privkey := NewChallengerDidDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetChallengerDidDocument(ctx, Did, didDocument1)

	didDocument2, privkey := NewChallengerDidDocumentWithSeq(SecondDid)
	require.NotNil(t, privkey)
	keeper.SetChallengerDidDocument(ctx, SecondDid, didDocument2)

	// Test first DidDocument
	got, found := keeper.GetChallengerDidDocument(ctx, Did)
	require.Equal(t, true, found)
	require.NotNil(t, got)
	require.Equal(t, didDocument1, got)

	// Test all DIDs
	resDids := keeper.GetAllChallengerDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}
