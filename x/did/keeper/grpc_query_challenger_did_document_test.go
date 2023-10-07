package keeper_test

import (
	"soarchain/x/did/types"
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func Test_GetChallengerDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	didDocument, privkey := NewChallengerDidDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetChallengerDidDocument(ctx, Did, didDocument)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerDidRequest
		response *types.QueryGetChallengerDidResponse
		err      error
	}{
		{
			desc: "Valid Did Id",
			request: &types.QueryGetChallengerDidRequest{
				Id: didDocument.Document.Id,
			},
			response: &types.QueryGetChallengerDidResponse{ChallengerDidDocument: didDocument},
		},
		{
			desc: "Not Valid Did Id",
			request: &types.QueryGetChallengerDidRequest{
				Id: "Not-Valid",
			},
			response: &types.QueryGetChallengerDidResponse{ChallengerDidDocument: didDocument},
		},
		{
			desc: "Empty",
			request: &types.QueryGetChallengerDidRequest{
				Id: "",
			},
			response: &types.QueryGetChallengerDidResponse{ChallengerDidDocument: didDocument},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, found := keeper.GetChallengerDidDocument(ctx, Did)
			if !found {
				require.Equal(t, found, false)
				require.Nil(t, response)
			}
			require.Equal(t, found, true)
			require.NotNil(t, response)
		})
	}
}
