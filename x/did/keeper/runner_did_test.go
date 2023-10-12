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
	keeper.SetRunnerDid(ctx, *didDocument1.Document)

	got1, found1 := keeper.GetRunnerDid(ctx, ADDRESS)
	require.Equal(t, true, found1)
	require.NotNil(t, got1)

	didDocument2, privkey := NewRunnerDidDocumentWithSeq(SecondDid)
	didDocument2.Document.Address = SecondAddress
	didDocument2.Document.PubKey = SecondPubKey
	require.NotNil(t, privkey)
	keeper.SetRunnerDid(ctx, *didDocument2.Document)

	got2, found2 := keeper.GetRunnerDid(ctx, ADDRESS)
	require.Equal(t, true, found2)
	require.NotNil(t, got2)

	// Test all DIDs
	resDids := keeper.GetAllRunnerDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}
