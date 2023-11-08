package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SetGetChallengerDid(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	challengers := CreateNChallengerDid(keeper, ctx, 1)
	targetChallenger := challengers[0]

	res, found := keeper.GetChallengerDid(ctx, targetChallenger.Address)
	require.Equal(t, true, found)
	require.Equal(t, res, targetChallenger)

	challengersList := keeper.GetAllChallengerDid(ctx)
	require.NotNil(t, challengersList)
	require.Contains(t, challengersList, targetChallenger)
}

// func Test_GetChallengerDidByPubkey(t *testing.T) {
// 	keeper, ctx := keepertest.DidKeeper(t)

// 	didDocument, privkey := NewChallengerDidDocumentWithSeq(Did)
// 	require.NotNil(t, privkey)
// 	keeper.SetChallengerDid(ctx, *didDocument.Document)

// 	got, found := keeper.GetChallengerDidUsingPubKey(ctx, didDocument.Document.PubKey)
// 	t.Log("challenger did -->", got, found)
// 	require.Equal(t, true, found)
// 	require.NotNil(t, got)

// }
