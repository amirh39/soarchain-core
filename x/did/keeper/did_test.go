package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func Test_SetGetDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	didDocument1, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, Did, didDocument1)

	didDocument2, privkey := NewDIDDocumentWithSeq(SecondDid)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, SecondDid, didDocument2)

	// Test first DidDocument
	got, found := keeper.GetDidDocument(ctx, Did)
	require.Equal(t, true, found)
	require.NotNil(t, got)
	require.Equal(t, didDocument1, got)

	// Test all DIDs
	resDids := keeper.GetAllDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}
