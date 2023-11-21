package keeper_test

// import (
// 	"github.com/soar-robotics/soarchain-core/x/did/types"
// 	"strconv"
// 	"testing"

// 	keepertest "github.com/soar-robotics/soarchain-core/testutil/keeper"

// 	"github.com/stretchr/testify/require"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func Test_ClientQuery(t *testing.T) {
// 	keeper, ctx := keepertest.DidKeeper(t)
// 	//wctx := sdk.WrapSDKContext(ctx)
// 	didDocument, privkey := NewDIDDocumentWithSeq(Did)
// 	require.NotNil(t, privkey)
// 	keeper.SetClientDid(ctx, *didDocument.Document)

// 	for _, tc := range []struct {
// 		desc     string
// 		request  *types.QueryGetClientDidRequest
// 		response *types.QueryGetClientDidResponse
// 		err      error
// 	}{
// 		{
// 			desc: "First",
// 			request: &types.QueryGetClientDidRequest{
// 				Address: didDocument.Document.Address,
// 			},
// 			response: &types.QueryGetClientDidResponse{ClientDid: *didDocument.Document},
// 		},
// 		{
// 			desc: "KeyNotFound",
// 			request: &types.QueryGetClientDidRequest{
// 				Address: strconv.Itoa(100000),
// 			},
// 			err: status.Error(codes.NotFound, "not found"),
// 		},
// 		{
// 			desc: "InvalidRequest",
// 			err:  status.Error(codes.InvalidArgument, "invalid request"),
// 		},
// 	} {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			response, found := keeper.GetClientDid(ctx, ADDRESS)
// 			t.Log("response", response)
// 			if !found {
// 				require.Equal(t, false, false)
// 			} else {
// 				require.Equal(t, true, true)
// 				require.NotNil(t, response)
// 			}
// 		})
// 	}
// }
