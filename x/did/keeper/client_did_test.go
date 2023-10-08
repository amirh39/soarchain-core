package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SetGetDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	didDocument1, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetClientDid(ctx, *didDocument1.Document)

	didDocument2, privkey := NewDIDDocumentWithSeq(SecondDid)
	require.NotNil(t, privkey)
	keeper.SetClientDid(ctx, *didDocument2.Document)

	// Test first DidDocument
	got, found := keeper.GetClientDid(ctx, ADDRESS)
	require.Equal(t, true, found)
	require.NotNil(t, got)
	require.Equal(t, didDocument1, got)

	// Test all DIDs
	resDids := keeper.GetAllClientDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}
