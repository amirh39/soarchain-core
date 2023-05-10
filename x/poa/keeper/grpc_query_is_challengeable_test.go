
package keeper_test

import (
	"testing"

	"soarchain/x/poa/types"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"strconv"

	keepertest "soarchain/testutil/keeper"
	"github.com/stretchr/testify/require"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createValidClient(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(12)
		items[i].Score = "80"
		items[i].LastTimeChallenged = "2023-01-06 11:05:17.40125 +0000 UTC"
		items[i].CoolDownTolerance = "1"
		keeper.SetClient(ctx, items[i])
	}
	return items
}

/** Test challengibility with valid client entity */
func Test_IsChallengeable(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	clients := createValidClient(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryIsChallengeableRequest
		response  *types.QueryIsChallengeableResponse
		err      error
	}{
		{
			desc: "Valid Client Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: clients[0].Index,
			},
			response: &types.QueryIsChallengeableResponse{ResultBool: "false", ChallengeabilityScore: "-63808599897"},
		},
	}{
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.IsChallengeable(wctx, tc.request)
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

/** If one of client entity parameters are invalid challengibility would not be calculated. */
func createNotValidClient(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(12)
		items[i].Score = ""
		items[i].LastTimeChallenged = "-01-06 11:05:17.40125 +0000 UTC"
		items[i].CoolDownTolerance = "10"
		keeper.SetClient(ctx, items[i])
	}
	return items
}

/** Test challengibility with not valid client entity. All tests should crash and proper error message will raise. */
func Test_IsNotChallengeable(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	clients := createNotValidClient(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryIsChallengeableRequest
		response *types.QueryIsChallengeableResponse
		err      error
	}{
		{
			desc: "Invalid Client Address",
			request: &types.QueryIsChallengeableRequest{
				ClientAddr: clients[0].Index,
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
	}{
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.IsChallengeable(wctx, tc.request)
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
	