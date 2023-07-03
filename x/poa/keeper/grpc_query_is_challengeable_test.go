package keeper_test

import (
	"testing"

	"soarchain/testutil/nullify"
	"soarchain/x/poa/types"
	"strconv"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/** Test challengibility with valid client entity */
func Test_IsChallengeable(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	clients := CreateNClient(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryIsChallengeableRequest
		response *types.QueryIsChallengeableResponse
		err      error
	}{
		{
			desc: "Valid Client Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: clients[0].Address,
			},
			response: &types.QueryIsChallengeableResponse{ResultBool: "false", ChallengeabilityScore: "-63808599897"},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.IsChallengeable(wctx, tc.request)
			if err != nil {
				require.Error(t, err)
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

/** If one of client entity parameters are invalid challengibility would not be calculated. */

/** Test challengibility with not valid client entity. All tests should crash and proper error message will raise. */
func Test_IsNotChallengeable(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	clients := CreateInValidClientScore(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryIsChallengeableRequest
		response *types.QueryIsChallengeableResponse
		err      error
	}{
		{
			desc: "Invalid Client Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: clients[0].Address,
			},
			response: &types.QueryIsChallengeableResponse{ResultBool: "false", ChallengeabilityScore: "-63808599897"},
		},
		{
			desc: "Invalid Client Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: "0xc000db6e50",
			},
			err: status.Error(codes.NotFound, "client not found"),
		},
		{
			desc: "Client Address Not Found",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "client not found"),
		},
		{
			desc: "Invalid Request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.IsChallengeable(wctx, tc.request)
			if err != nil {
				require.Error(t, err)
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
