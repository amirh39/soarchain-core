package keeper_test

import (
	"soarchain/x/did/types"
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func Test_GetDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	didDocument, privkey := NewDIDDocumentWithSeq(Did)
	require.NotNil(t, privkey)
	keeper.SetDidDocument(ctx, Did, didDocument)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDidRequest
		response *types.QueryGetDidResponse
		err      error
	}{
		{
			desc: "Valid Did Id",
			request: &types.QueryGetDidRequest{
				Id: didDocument.Document.Id,
			},
			response: &types.QueryGetDidResponse{DidDocument: didDocument},
		},
		{
			desc: "Not Valid Did Id",
			request: &types.QueryGetDidRequest{
				Id: "Not-Valid",
			},
			response: &types.QueryGetDidResponse{DidDocument: didDocument},
		},
		{
			desc: "Empty",
			request: &types.QueryGetDidRequest{
				Id: "",
			},
			response: &types.QueryGetDidResponse{DidDocument: didDocument},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, found := keeper.GetDidDocument(ctx, Did)
			if !found {
				require.Equal(t, found, false)
				require.Nil(t, response)
			}
			require.Equal(t, found, true)
			require.NotNil(t, response)
		})
	}
}
