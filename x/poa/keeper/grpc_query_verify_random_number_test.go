package keeper_test

import (
	"testing"

	"soarchain/testutil/nullify"
	"soarchain/x/poa/types"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_VerifyRandomNumber(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryVerifyRandomNumberRequest
		response *types.QueryVerifyRandomNumberResponse
		err      error
	}{
		{
			desc: "Valid Random Number",
			request: &types.QueryVerifyRandomNumberRequest{
				Pubkey:  ChallengerPubKey,
				Message: RandomNumber_Message,
				Vrv:     RandomNumber_Vrv,
				Proof:   RandomNumber_Proof,
			},
			response: &types.QueryVerifyRandomNumberResponse{Result: false},
		},
	} {
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
