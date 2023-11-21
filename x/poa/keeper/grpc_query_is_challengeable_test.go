package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/soar-robotics/soarchain-core/testutil/keeper"
	"github.com/soar-robotics/soarchain-core/testutil/nullify"
	"github.com/soar-robotics/soarchain-core/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/** Test challengibility with valid reputation entity */
func Test_IsChallengeable(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	reputations := CreateNReputation(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryIsChallengeableRequest
		response *types.QueryIsChallengeableResponse
		err      error
	}{
		{
			desc: "Valid reputation Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: reputations[0].Address,
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

/** If one of reputation entity parameters are invalid challengibility would not be calculated. */
func Test_IsNotChallengeable(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	reputations := CreateInValidReputationScore(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryIsChallengeableRequest
		response *types.QueryIsChallengeableResponse
		err      error
	}{
		{
			desc: "Invalid reputation Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: reputations[0].Address,
			},
			response: &types.QueryIsChallengeableResponse{ResultBool: "false", ChallengeabilityScore: "-63808599897"},
		},
		{
			desc: "Invalid reputation Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: "0xc000db6e50",
			},
			err: status.Error(codes.NotFound, "reputation not found"),
		},
		{
			desc: "reputation Address Not Found",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "reputation not found"),
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
