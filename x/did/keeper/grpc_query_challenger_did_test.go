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
	keeper.SetChallengerDid(ctx, *didDocument.Document)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerDidRequest
		response *types.QueryGetChallengerDidResponse
		err      error
	}{
		{
			desc: "Valid Did Id",
			request: &types.QueryGetChallengerDidRequest{
				Address: ADDRESS,
			},
			response: &types.QueryGetChallengerDidResponse{ChallengerDid: *didDocument.Document},
		},
		{
			desc: "Not Valid Did Id",
			request: &types.QueryGetChallengerDidRequest{
				Address: "Not-Valid",
			},
			response: &types.QueryGetChallengerDidResponse{ChallengerDid: *didDocument.Document},
		},
		{
			desc: "Empty",
			request: &types.QueryGetChallengerDidRequest{
				Address: "",
			},
			response: &types.QueryGetChallengerDidResponse{ChallengerDid: *didDocument.Document},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, found := keeper.GetChallengerDid(ctx, ADDRESS)
			if !found {
				require.Equal(t, found, false)
				require.Nil(t, response)
			}
			require.Equal(t, found, true)
			require.NotNil(t, response)
		})
	}
}
