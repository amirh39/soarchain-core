package keeper_test

import (
	"soarchain/x/did/types"
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func Test_GetRunnerDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	didDocument, privkey := NewRunnerDidDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetRunnerDidDocument(ctx, Did, didDocument)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRunnerDidRequest
		response *types.QueryGetRunnerDidResponse
		err      error
	}{
		{
			desc: "Valid Did Id",
			request: &types.QueryGetRunnerDidRequest{
				Id: didDocument.Document.Id,
			},
			response: &types.QueryGetRunnerDidResponse{RunnerDidDocument: didDocument},
		},
		{
			desc: "Not Valid Did Id",
			request: &types.QueryGetRunnerDidRequest{
				Id: "Not-Valid",
			},
			response: &types.QueryGetRunnerDidResponse{RunnerDidDocument: didDocument},
		},
		{
			desc: "Empty",
			request: &types.QueryGetRunnerDidRequest{
				Id: "",
			},
			response: &types.QueryGetRunnerDidResponse{RunnerDidDocument: didDocument},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, found := keeper.GetRunnerDidDocument(ctx, Did)
			if !found {
				require.Equal(t, found, false)
				require.Nil(t, response)
			}
			require.Equal(t, found, true)
			require.NotNil(t, response)
		})
	}
}
