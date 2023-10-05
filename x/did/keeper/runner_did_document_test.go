package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SetGetRunnerDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	didDocument1, privkey := NewRunnerDidDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetRunnerDidDocument(ctx, Did, didDocument1)

	didDocument2, privkey := NewRunnerDidDocumentWithSeq(SecondDid)
	require.NotNil(t, privkey)
	keeper.SetRunnerDidDocument(ctx, SecondDid, didDocument2)

	// Test first DidDocument
	got, found := keeper.GetRunnerDidDocument(ctx, Did)
	require.Equal(t, true, found)
	require.NotNil(t, got)
	require.Equal(t, didDocument1, got)

	// Test all DIDs
	resDids := keeper.GetAllRunnerDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}
