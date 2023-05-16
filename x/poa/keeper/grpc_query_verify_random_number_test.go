
package keeper_test

import (
	"testing"

	"soarchain/x/poa/types"
	"soarchain/testutil/nullify"

	keepertest "soarchain/testutil/keeper"
	"github.com/stretchr/testify/require"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Test_VerifyRandomNumber(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryVerifyRandomNumberRequest
		response  *types.QueryVerifyRandomNumberResponse
		err      error
	}{
		{
			desc: "Valid Random Number",
			request: &types.QueryVerifyRandomNumberRequest{
				Pubkey: "3056301006072a8648ce3d020106052b8104000a03420004c4039cc2459a57357707620ddbbaddfeda5d4c66cc9ac9c3aac997e65f16b78253b3f9241182014246c1b945595c1ed2463e22ca59f153a74fee375e23a86561",
				Message: "message",
				Vrv: "3",
				Proof: "3",
			},
			response: &types.QueryVerifyRandomNumberResponse{Result: false},
		},
	}{
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VerifyRandomNumber(wctx, tc.request)
			if tc.err != nil {
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