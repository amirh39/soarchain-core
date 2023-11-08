package keeper_test

// import (
// 	"soarchain/x/did/types"
// 	"testing"

// 	keepertest "soarchain/testutil/keeper"

// 	"github.com/stretchr/testify/require"
// )

// func Test_GetRunnerDidDocument(t *testing.T) {
// 	keeper, ctx := keepertest.DidKeeper(t)
// 	didDocument, privkey := NewRunnerDidDocumentWithSeq(Did)
// 	require.NotNil(t, privkey)
// 	keeper.SetRunnerDid(ctx, *didDocument.Document)

// 	for _, tc := range []struct {
// 		desc     string
// 		request  *types.QueryGetRunnerDidRequest
// 		response *types.QueryGetRunnerDidResponse
// 		err      error
// 	}{
// 		{
// 			desc: "Valid Did Id",
// 			request: &types.QueryGetRunnerDidRequest{
// 				Address: didDocument.Document.Id,
// 			},
// 			response: &types.QueryGetRunnerDidResponse{RunnerDid: *didDocument.Document},
// 		},
// 		{
// 			desc: "Not Valid Did Id",
// 			request: &types.QueryGetRunnerDidRequest{
// 				Address: "Not-Valid",
// 			},
// 			response: &types.QueryGetRunnerDidResponse{RunnerDid: *didDocument.Document},
// 		},
// 		{
// 			desc: "Empty",
// 			request: &types.QueryGetRunnerDidRequest{
// 				Address: "",
// 			},
// 			response: &types.QueryGetRunnerDidResponse{RunnerDid: *didDocument.Document},
// 		},
// 	} {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			response, found := keeper.GetRunnerDid(ctx, ADDRESS)
// 			if !found {
// 				require.Equal(t, found, false)
// 				require.Nil(t, response)
// 			}
// 			require.Equal(t, found, true)
// 			require.NotNil(t, response)
// 		})
// 	}
// }
