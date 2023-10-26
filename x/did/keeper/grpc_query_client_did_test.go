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
	keeper.SetClientDid(ctx, *didDocument.Document)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetClientDidRequest
		response *types.QueryGetClientDidResponse
		err      error
	}{
		{
			desc: "Valid Did Id",
			request: &types.QueryGetClientDidRequest{
				Address: didDocument.Document.Address,
			},
			response: &types.QueryGetClientDidResponse{ClientDid: *didDocument.Document},
		},
		{
			desc: "Not Valid Did Id",
			request: &types.QueryGetClientDidRequest{
				Address: "Not-Valid",
			},
			response: &types.QueryGetClientDidResponse{ClientDid: *didDocument.Document},
		},
		{
			desc: "Empty",
			request: &types.QueryGetClientDidRequest{
				Address: "",
			},
			response: &types.QueryGetClientDidResponse{ClientDid: *didDocument.Document},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, found := keeper.GetClientDid(ctx, ADDRESS)
			if !found {
				require.Equal(t, found, false)
				require.Nil(t, response)
			}
			require.Equal(t, found, true)
			require.NotNil(t, response)
		})
	}
}
