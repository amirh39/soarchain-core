package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
)

func Test_DidDocumentGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	didDocument1, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, Did, didDocument1)

	didDocument2, privkey := NewDIDDocumentWithSeq(SecondDid)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, SecondDid, didDocument2)

	got, found := keeper.GetDidDocument(ctx, Did)
	t.Log("Fetched Did", got)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&didDocument1),
		nullify.Fill(&got),
	)
}

func Test_DidsDocumentGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	didDocument1, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, Did, didDocument1)

	didDocument2, privkey := NewDIDDocumentWithSeq(SecondDid)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, SecondDid, didDocument2)

	// Test all DIDs
	resDids := keeper.GetAllDid(ctx)
	t.Log("all dids", resDids)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}
