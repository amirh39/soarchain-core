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
	keeper.SetClientDidDocument(ctx, Did, clientDidDocument)

	isFound := keeper.IsUniqueDid(ctx, clientDidDocument.Document.Id)
	require.Equal(t, true, isFound)
}
