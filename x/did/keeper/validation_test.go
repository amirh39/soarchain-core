package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IsUniqueDid(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument
	clientDidDocument, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetClientDid(ctx, *clientDidDocument.Document)

	isFound := keeper.IsNotUniqueDid(ctx, clientDidDocument.Document.Id)
	require.Equal(t, true, isFound)
}
