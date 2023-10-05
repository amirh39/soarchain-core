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
	keeper.SetClientDidDocument(ctx, Did, didDocument)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetClientDidRequest
		response *types.QueryGetClientDidResponse
		err      error
	}{
		{
			desc: "Valid Did Id",
			request: &types.QueryGetClientDidRequest{
				Id: didDocument.Document.Id,
			},
			response: &types.QueryGetClientDidResponse{ClientDidDocument: didDocument},
		},
		{
			desc: "Not Valid Did Id",
			request: &types.QueryGetClientDidRequest{
				Id: "Not-Valid",
			},
			response: &types.QueryGetClientDidResponse{ClientDidDocument: didDocument},
		},
		{
			desc: "Empty",
			request: &types.QueryGetClientDidRequest{
				Id: "",
			},
			response: &types.QueryGetClientDidResponse{ClientDidDocument: didDocument},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, found := keeper.GetClientDidDocument(ctx, Did)
			if !found {
				require.Equal(t, found, false)
				require.Nil(t, response)
			}
			require.Equal(t, found, true)
			require.NotNil(t, response)
		})
	}
}
