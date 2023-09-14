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
	got, found := keeper.GetDidDocumentWithSequence(ctx, Did)
	require.Equal(t, true, found)
	require.NotNil(t, got)
	require.Equal(t, didDocument1, got)

	// Test all DIDs
	resDids := keeper.GetAllDid(ctx)
	require.NotNil(t, resDids)
	require.Equal(t, 2, len(resDids))
}

func Test_GetDidDocumentByPubkey(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	didDocument1, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, Did, didDocument1)
	require.NotNil(t, didDocument1)

	dids, found := keeper.GetDidDocumentByPubkey(ctx, PUBKEY)
	if !found {
		require.Equal(t, found, false)
	} else {
		require.NotNil(t, dids)
	}
}

func Test_GetDidDocumentByPin(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	didDocument1, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, Did, didDocument1)
	require.NotNil(t, didDocument1)
	var selectedPin = []uint{1}
	found := keeper.FindEligibleDid(ctx, selectedPin)
	t.Log("kkkkkkkkkkkkkkkkkk-found", found)
}
