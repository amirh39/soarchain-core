package keeper_test

import (
	"testing"

	"soarchain/testutil/nullify"
	"soarchain/x/did/types"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/** The public key which is received by the request is a valid Pubkey. Response should return a valid challenger
which is related to that public key.*/
func Test_GetChallengerByPubKey(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNChallengerDid(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerDidByPubKeyRequest
		response *types.QueryGetChallengerDidByPubKeyResponse
		err      error
	}{
		{
			desc: "Valid Public Key",
			request: &types.QueryGetChallengerDidByPubKeyRequest{
				Pubkey: msgs[0].PubKey,
			},
			response: &types.QueryGetChallengerDidByPubKeyResponse{ChallengerDid: &msgs[0]},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetChallengerDidByPubKey(wctx, tc.request)

			if err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
